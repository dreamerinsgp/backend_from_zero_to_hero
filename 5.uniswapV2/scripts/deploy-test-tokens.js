/**
 * Deploy two instances of contracts/test/ERC20.sol on the same network (e.g. Sepolia).
 * Same .env as scripts/deploy.js: RPC_URL, PRIVATE_KEY
 *
 * Optional env:
 *   SUPPLY_A_WEI  — total supply for token A (default: 1e24 wei = 1e6 tokens * 1e18)
 *   SUPPLY_B_WEI  — total supply for token B (default: same as A)
 *
 * Usage: yarn deploy:tokens
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

function loadArtifact(name) {
  const artifactPath = path.join(__dirname, '..', 'build', `${name}.json`)
  if (!fs.existsSync(artifactPath)) {
    throw new Error(`Missing ${artifactPath} — run yarn compile first`)
  }
  return JSON.parse(fs.readFileSync(artifactPath, 'utf8'))
}

function getBytecode(artifact) {
  if (artifact.bytecode && typeof artifact.bytecode === 'string') {
    return artifact.bytecode.startsWith('0x') ? artifact.bytecode : '0x' + artifact.bytecode
  }
  const obj = artifact.evm && artifact.evm.bytecode && artifact.evm.bytecode.object
  if (obj) return obj.startsWith('0x') ? obj : '0x' + obj
  throw new Error('bytecode not found in artifact')
}

async function main() {
  const rpcUrl = process.env.RPC_URL
  if (!rpcUrl) throw new Error('Set RPC_URL in .env')

  const provider = new ethers.providers.JsonRpcProvider(rpcUrl)
  const wallet = new ethers.Wallet(normalizePrivateKey(process.env.PRIVATE_KEY), provider)
  console.log('Deployer:', wallet.address)

  const net = await provider.getNetwork()
  console.log('chainId:', net.chainId.toString())

  const defaultSupply = ethers.utils.parseEther('1000000') // 1e6 tokens, 18 decimals
  const supplyA = process.env.SUPPLY_A_WEI
    ? ethers.BigNumber.from(process.env.SUPPLY_A_WEI)
    : defaultSupply
  const supplyB = process.env.SUPPLY_B_WEI
    ? ethers.BigNumber.from(process.env.SUPPLY_B_WEI)
    : defaultSupply

  const artifact = loadArtifact('ERC20')
  const bytecode = getBytecode(artifact)
  const Factory = new ethers.ContractFactory(artifact.abi, bytecode, wallet)

  console.log('Deploying token A, supply (wei):', supplyA.toString())
  const tokenA = await Factory.deploy(supplyA)
  await tokenA.deployed()
  console.log('TOKEN_A=', tokenA.address)

  console.log('Deploying token B, supply (wei):', supplyB.toString())
  const tokenB = await Factory.deploy(supplyB)
  await tokenB.deployed()
  console.log('TOKEN_B=', tokenB.address)

  console.log('\nUse with Factory.createPair (order does not matter):')
  console.log(`  createPair(${tokenA.address}, ${tokenB.address})`)
}

main().catch((e) => {
  console.error(e.message || e)
  process.exit(1)
})
