/**
 * Call UniswapV2Factory.createPair(tokenA, tokenB) on a live network (e.g. Sepolia).
 *
 * Prerequisites: yarn compile (for ABI), .env with RPC_URL + PRIVATE_KEY
 *
 * Required env:
 *   FACTORY_ADDRESS — deployed UniswapV2Factory
 *   TOKEN_A         — first ERC20 address
 *   TOKEN_B         — second ERC20 address
 *
 * Usage: yarn createpair
 */
require('dotenv').config()

const fs = require('fs')
const path = require('path')
const { ethers } = require('ethers')

function normalizePrivateKey(pk) {
  if (!pk || typeof pk !== 'string') {
    throw new Error('Set PRIVATE_KEY in .env')
  }
  const t = pk.trim()
  return t.startsWith('0x') ? t : '0x' + t
}

function loadFactoryAbi() {
  const artifactPath = path.join(__dirname, '..', 'build', 'UniswapV2Factory.json')
  if (!fs.existsSync(artifactPath)) {
    throw new Error('Missing build/UniswapV2Factory.json — run yarn compile first')
  }
  const artifact = JSON.parse(fs.readFileSync(artifactPath, 'utf8'))
  return artifact.abi
}

function reqAddr(name) {
  const v = process.env[name]
  if (!v || !ethers.utils.isAddress(v)) {
    throw new Error(`Set ${name} in .env to a valid checksummed address`)
  }
  return ethers.utils.getAddress(v.trim())
}

async function main() {
  const rpcUrl = process.env.RPC_URL
  if (!rpcUrl) throw new Error('Set RPC_URL in .env')

  const factoryAddr = reqAddr('FACTORY_ADDRESS')
  const tokenA = reqAddr('TOKEN_A')
  const tokenB = reqAddr('TOKEN_B')

  if (tokenA === tokenB) {
    throw new Error('TOKEN_A and TOKEN_B must be different')
  }

  const provider = new ethers.providers.JsonRpcProvider(rpcUrl)
  const net = await provider.getNetwork()
  console.log('chainId:', net.chainId.toString())

  const wallet = new ethers.Wallet(normalizePrivateKey(process.env.PRIVATE_KEY), provider)
  console.log('Caller:', wallet.address)

  const abi = loadFactoryAbi()
  const factory = new ethers.Contract(factoryAddr, abi, wallet)

  const existing = await factory.getPair(tokenA, tokenB)
  if (existing !== ethers.constants.AddressZero) {
    console.log('Pair already exists:', existing)
    console.log('Nothing to do.')
    return
  }

  console.log('Sending createPair...')
  const tx = await factory.createPair(tokenA, tokenB)
  console.log('Tx hash:', tx.hash)
  const receipt = await tx.wait()
  console.log('Confirmed in block:', receipt.blockNumber)

  const pair = await factory.getPair(tokenA, tokenB)
  console.log('PAIR_ADDRESS=', pair)

  for (const log of receipt.logs) {
    try {
      const parsed = factory.interface.parseLog(log)
      if (parsed.name === 'PairCreated') {
        console.log(
          'PairCreated — token0:',
          parsed.args.token0,
          'token1:',
          parsed.args.token1,
          'pair:',
          parsed.args.pair
        )
        break
      }
    } catch (_) {
      /* not this contract's event */
    }
  }
}

main().catch((e) => {
  console.error(e.message || e)
  process.exit(1)
})
