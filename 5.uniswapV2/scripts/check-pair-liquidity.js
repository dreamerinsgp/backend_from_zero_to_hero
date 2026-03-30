/**
 * Read-only: print Uniswap V2 Pair reserves, LP total supply, and token metadata.
 *
 * Prerequisites: yarn compile (for Pair ABI)
 *
 * Required env:
 *   RPC_URL
 *   PAIR_ADDRESS
 *
 * Usage: yarn check:liquidity
 */
require('dotenv').config()

const fs = require('fs')
const path = require('path')
const { ethers } = require('ethers')

const ERC20_META = [
  'function symbol() view returns (string)',
  'function decimals() view returns (uint8)',
]

function loadPairAbi() {
  const artifactPath = path.join(__dirname, '..', 'build', 'UniswapV2Pair.json')
  if (!fs.existsSync(artifactPath)) {
    throw new Error('Missing build/UniswapV2Pair.json — run yarn compile first')
  }
  const artifact = JSON.parse(fs.readFileSync(artifactPath, 'utf8'))
  return artifact.abi
}

function reqAddr(name) {
  const v = process.env[name]
  if (!v || !ethers.utils.isAddress(v)) {
    throw new Error(`Set ${name} in .env to a valid address`)
  }
  return ethers.utils.getAddress(v.trim())
}

async function tokenMeta(provider, tokenAddr) {
  const c = new ethers.Contract(tokenAddr, ERC20_META, provider)
  let symbol = '?'
  let decimals = 18
  try {
    symbol = await c.symbol()
  } catch (_) {
    /* non-standard */
  }
  try {
    decimals = await c.decimals()
  } catch (_) {
    decimals = 18
  }
  return { symbol, decimals }
}

async function main() {
  const rpcUrl = process.env.RPC_URL
  if (!rpcUrl) throw new Error('Set RPC_URL in .env')

  const pairAddr = reqAddr('PAIR_ADDRESS')
  const provider = new ethers.providers.JsonRpcProvider(rpcUrl)
  const net = await provider.getNetwork()
  console.log('chainId:', net.chainId.toString())
  console.log('PAIR_ADDRESS:', pairAddr)

  const abi = loadPairAbi()
  const pair = new ethers.Contract(pairAddr, abi, provider)

  const [token0, token1] = await Promise.all([pair.token0(), pair.token1()])
  const [meta0, meta1] = await Promise.all([tokenMeta(provider, token0), tokenMeta(provider, token1)])

  const [reserve0, reserve1, blockTimestampLast] = await pair.getReserves()
  const totalSupply = await pair.totalSupply()

  console.log('')
  console.log('token0:', token0, `(${meta0.symbol})`)
  console.log('token1:', token1, `(${meta1.symbol})`)
  console.log('')
  console.log('reserve0 (raw):', reserve0.toString())
  console.log('reserve1 (raw):', reserve1.toString())
  console.log(
    'reserves (formatted):',
    ethers.utils.formatUnits(reserve0, meta0.decimals),
    meta0.symbol,
    '/',
    ethers.utils.formatUnits(reserve1, meta1.decimals),
    meta1.symbol
  )
  console.log('blockTimestampLast:', blockTimestampLast.toString())
  console.log('')
  console.log('LP totalSupply (raw):', totalSupply.toString())
  console.log('LP totalSupply (wei as decimal):', ethers.utils.formatUnits(totalSupply, 18))
}

main().catch((e) => {
  console.error(e.message || e)
  process.exit(1)
})
