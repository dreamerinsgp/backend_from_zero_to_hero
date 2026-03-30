/**
 * Add liquidity to a Uniswap V2 Pair: transfer token0 + token1 to the pair, then mint(to).
 *
 * Prerequisites: yarn compile (Pair ABI), .env with RPC_URL + PRIVATE_KEY
 *
 * Required env:
 *   PAIR_ADDRESS
 *   Either:
 *     AMOUNT0_WEI + AMOUNT1_WEI  — amounts in wei for Pair.token0() / Pair.token1()
 *   Or (friendlier if you only know TOKEN_A / TOKEN_B from deploy scripts):
 *     TOKEN_A + TOKEN_B + AMOUNT_A_WEI + AMOUNT_B_WEI
 *     (amounts apply to those token addresses; script maps to token0/token1)
 *
 * Optional:
 *   LP_TO — address that receives LP tokens (default: caller)
 *
 * Usage: yarn add:liquidity
 *
 * Note: three sequential txs (transfer, transfer, mint). For production use a Router
 * (v2-periphery) for atomic addLiquidity + slippage. See questions/6.add_liquidity.md
 */
require('dotenv').config()

const fs = require('fs')
const path = require('path')
const { ethers } = require('ethers')

const ERC20_MIN = [
  'function transfer(address to, uint256 amount) returns (bool)',
  'function balanceOf(address owner) view returns (uint256)',
]

function normalizePrivateKey(pk) {
  if (!pk || typeof pk !== 'string') {
    throw new Error('Set PRIVATE_KEY in .env')
  }
  const t = pk.trim()
  return t.startsWith('0x') ? t : '0x' + t
}

function loadPairAbi() {
  const artifactPath = path.join(__dirname, '..', 'build', 'UniswapV2Pair.json')
  if (!fs.existsSync(artifactPath)) {
    throw new Error('Missing build/UniswapV2Pair.json — run yarn compile first')
  }
  return JSON.parse(fs.readFileSync(artifactPath, 'utf8')).abi
}

function reqAddr(name) {
  const v = process.env[name]
  if (!v || !ethers.utils.isAddress(v)) {
    throw new Error(`Set ${name} in .env to a valid address`)
  }
  return ethers.utils.getAddress(v.trim())
}

function reqPositiveBn(name) {
  const v = process.env[name]
  if (v === undefined || v === '') {
    throw new Error(`Set ${name} in .env (wei, decimal integer string)`)
  }
  const bn = ethers.BigNumber.from(v.trim())
  if (bn.lte(0)) throw new Error(`${name} must be > 0`)
  return bn
}

function throwMissingAmounts(token0, token1) {
  const msg = [
    'Missing liquidity amounts in .env. Add one of the following:',
    '',
    'Option A — amounts for Pair.token0 / token1 (wei, integer string):',
    `  AMOUNT0_WEI=1000000000000000000`,
    `  AMOUNT1_WEI=1000000000000000000`,
    `  (token0: ${token0})`,
    `  (token1: ${token1})`,
    '',
    'Option B — your deploy labels (must match this pair):',
    `  TOKEN_A=${token0}`,
    `  TOKEN_B=${token1}`,
    '  AMOUNT_A_WEI=1000000000000000000',
    '  AMOUNT_B_WEI=1000000000000000000',
    '',
    'Then run: yarn add:liquidity',
  ].join('\n')
  throw new Error(msg)
}

function resolveAmounts(token0, token1) {
  const hasAB =
    process.env.TOKEN_A &&
    process.env.TOKEN_B &&
    process.env.AMOUNT_A_WEI !== undefined &&
    process.env.AMOUNT_A_WEI !== '' &&
    process.env.AMOUNT_B_WEI !== undefined &&
    process.env.AMOUNT_B_WEI !== ''

  const has01 =
    process.env.AMOUNT0_WEI !== undefined &&
    process.env.AMOUNT0_WEI !== '' &&
    process.env.AMOUNT1_WEI !== undefined &&
    process.env.AMOUNT1_WEI !== ''

  if (hasAB && has01) {
    throw new Error('Use either TOKEN_A/B + AMOUNT_A/B or AMOUNT0/1_WEI, not both')
  }

  if (hasAB) {
    const a = reqAddr('TOKEN_A')
    const b = reqAddr('TOKEN_B')
    const set = new Set([a, b])
    if (set.size !== 2) throw new Error('TOKEN_A and TOKEN_B must differ')
    if (!set.has(token0) || !set.has(token1)) {
      throw new Error('TOKEN_A/TOKEN_B do not match this pair token0/token1')
    }
    const amtA = reqPositiveBn('AMOUNT_A_WEI')
    const amtB = reqPositiveBn('AMOUNT_B_WEI')
    if (a === token0) {
      return { amount0: amtA, amount1: amtB }
    }
    return { amount0: amtB, amount1: amtA }
  }

  if (has01) {
    return {
      amount0: reqPositiveBn('AMOUNT0_WEI'),
      amount1: reqPositiveBn('AMOUNT1_WEI'),
    }
  }

  throwMissingAmounts(token0, token1)
}

async function main() {
  const rpcUrl = process.env.RPC_URL
  if (!rpcUrl) throw new Error('Set RPC_URL in .env')

  const pairAddr = reqAddr('PAIR_ADDRESS')
  const provider = new ethers.providers.JsonRpcProvider(rpcUrl)
  const wallet = new ethers.Wallet(normalizePrivateKey(process.env.PRIVATE_KEY), provider)

  const net = await provider.getNetwork()
  console.log('chainId:', net.chainId.toString())
  console.log('Caller:', wallet.address)

  const pairAbi = loadPairAbi()
  const pair = new ethers.Contract(pairAddr, pairAbi, wallet)

  const token0 = await pair.token0()
  const token1 = await pair.token1()
  console.log('token0:', token0)
  console.log('token1:', token1)

  const { amount0, amount1 } = resolveAmounts(token0, token1)
  console.log('amount0 (wei):', amount0.toString())
  console.log('amount1 (wei):', amount1.toString())

  const erc0 = new ethers.Contract(token0, ERC20_MIN, wallet)
  const erc1 = new ethers.Contract(token1, ERC20_MIN, wallet)

  const [bal0, bal1] = await Promise.all([erc0.balanceOf(wallet.address), erc1.balanceOf(wallet.address)])
  if (bal0.lt(amount0)) throw new Error(`Insufficient token0 balance: have ${bal0.toString()}, need ${amount0.toString()}`)
  if (bal1.lt(amount1)) throw new Error(`Insufficient token1 balance: have ${bal1.toString()}, need ${amount1.toString()}`)

  const lpTo = process.env.LP_TO ? reqAddr('LP_TO') : wallet.address
  console.log('LP recipient:', lpTo)

  console.log('Transferring token0 to pair...')
  let tx = await erc0.transfer(pairAddr, amount0)
  console.log('  tx:', tx.hash)
  await tx.wait()

  console.log('Transferring token1 to pair...')
  tx = await erc1.transfer(pairAddr, amount1)
  console.log('  tx:', tx.hash)
  await tx.wait()

  console.log('Calling mint...')
  tx = await pair.mint(lpTo)
  console.log('  tx:', tx.hash)
  const receipt = await tx.wait()
  console.log('Confirmed in block:', receipt.blockNumber)

  const lpBal = await pair.balanceOf(lpTo)
  console.log('LP balance of recipient (raw):', lpBal.toString())
}

main().catch((e) => {
  console.error(e.message || e)
  process.exit(1)
})
