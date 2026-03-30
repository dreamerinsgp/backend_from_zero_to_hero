/**
 * Simple Uniswap V2 swap (exact input): transfer input token to Pair, then swap().
 *
 * Uses the same 0.3% fee formula as UniswapV2Library.getAmountOut (997/1000).
 *
 * Prerequisites: yarn compile, .env with RPC_URL + PRIVATE_KEY + PAIR_ADDRESS
 *
 * Required env:
 *   PAIR_ADDRESS
 *   SWAP_IN_TOKEN — "0" = sell token0 for token1, "1" = sell token1 for token0
 *   AMOUNT_IN_WEI — exact input amount (wei, integer string)
 *
 * Optional:
 *   SWAP_TO — recipient of output tokens (default: caller). Must NOT be token0 or token1 (INVALID_TO).
 *   MIN_OUT_WEI — revert if computed output < this (slippage guard)
 *
 * Usage: yarn swap
 *
 * Two transactions (transfer then swap). For production, use v2-periphery Router (atomic + path routing).
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

/**
 * Uniswap V2: amountIn * 997 * reserveOut / (reserveIn * 1000 + amountIn * 997)
 */
function getAmountOut(amountIn, reserveIn, reserveOut) {
  if (amountIn.lte(0)) throw new Error('amountIn must be > 0')
  if (reserveIn.lte(0) || reserveOut.lte(0)) throw new Error('Reserves must be positive — add liquidity first')
  const amountInWithFee = amountIn.mul(997)
  const numerator = amountInWithFee.mul(reserveOut)
  const denominator = reserveIn.mul(1000).add(amountInWithFee)
  return numerator.div(denominator)
}

async function main() {
  const rpcUrl = process.env.RPC_URL
  if (!rpcUrl) throw new Error('Set RPC_URL in .env')

  const pairAddr = reqAddr('PAIR_ADDRESS')
  const side = (process.env.SWAP_IN_TOKEN || '').trim()
  if (side !== '0' && side !== '1') {
    throw new Error('Set SWAP_IN_TOKEN to "0" (sell token0) or "1" (sell token1)')
  }

  const amountIn = (() => {
    const v = process.env.AMOUNT_IN_WEI
    if (v === undefined || v === '') throw new Error('Set AMOUNT_IN_WEI (wei integer string)')
    const bn = ethers.BigNumber.from(v.trim())
    if (bn.lte(0)) throw new Error('AMOUNT_IN_WEI must be > 0')
    return bn
  })()

  const provider = new ethers.providers.JsonRpcProvider(rpcUrl)
  const wallet = new ethers.Wallet(normalizePrivateKey(process.env.PRIVATE_KEY), provider)

  const net = await provider.getNetwork()
  console.log('chainId:', net.chainId.toString())
  console.log('Caller:', wallet.address)

  const pairAbi = loadPairAbi()
  const pair = new ethers.Contract(pairAddr, pairAbi, wallet)

  const token0 = await pair.token0()
  const token1 = await pair.token1()

  const swapTo = process.env.SWAP_TO ? reqAddr('SWAP_TO') : wallet.address
  if (swapTo === token0 || swapTo === token1) {
    throw new Error('SWAP_TO cannot be token0 or token1 (UniswapV2: INVALID_TO)')
  }

  const [reserve0, reserve1] = await pair.getReserves().then((r) => [r[0], r[1]])

  let amount0Out = ethers.BigNumber.from(0)
  let amount1Out = ethers.BigNumber.from(0)
  let tokenIn
  let amountOut

  if (side === '0') {
    tokenIn = new ethers.Contract(token0, ERC20_MIN, wallet)
    amountOut = getAmountOut(amountIn, reserve0, reserve1)
    amount1Out = amountOut
    console.log('Direction: token0 -> token1')
  } else {
    tokenIn = new ethers.Contract(token1, ERC20_MIN, wallet)
    amountOut = getAmountOut(amountIn, reserve1, reserve0)
    amount0Out = amountOut
    console.log('Direction: token1 -> token0')
  }

  console.log('Expected amount out (wei):', amountOut.toString())

  if (process.env.MIN_OUT_WEI !== undefined && process.env.MIN_OUT_WEI !== '') {
    const minOut = ethers.BigNumber.from(process.env.MIN_OUT_WEI.trim())
    if (amountOut.lt(minOut)) {
      throw new Error(`Output ${amountOut.toString()} < MIN_OUT_WEI ${minOut.toString()}`)
    }
  }

  const bal = await tokenIn.balanceOf(wallet.address)
  if (bal.lt(amountIn)) throw new Error(`Insufficient input token balance: have ${bal.toString()}, need ${amountIn.toString()}`)

  console.log('Transferring input to pair...')
  let tx = await tokenIn.transfer(pairAddr, amountIn)
  console.log('  tx:', tx.hash)
  await tx.wait()

  console.log('Calling swap...')
  tx = await pair.swap(amount0Out, amount1Out, swapTo, '0x')
  console.log('  tx:', tx.hash)
  const receipt = await tx.wait()
  console.log('Confirmed in block:', receipt.blockNumber)
}

main().catch((e) => {
  console.error(e.message || e)
  process.exit(1)
})
