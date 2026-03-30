/**
 * Deploy UniswapV2Factory to a network (e.g. Sepolia) using ethers v5.
 *
 * Prerequisites:
 *   yarn compile
 *   Copy .env.example to .env and fill RPC_URL + PRIVATE_KEY
 *
 * Usage: yarn deploy
 */
require('dotenv').config()

const fs = require('fs')
const path = require('path')
const { ethers } = require('ethers')

function normalizePrivateKey(pk) {
  if (!pk || typeof pk !== 'string') {
    throw new Error('Set PRIVATE_KEY in .env (hex string, with or without 0x)')
  }
  const t = pk.trim()
  if (t.startsWith('0x')) return t
  return '0x' + t
}

function loadArtifact() {
  const artifactPath = path.join(__dirname, '..', 'build', 'UniswapV2Factory.json')
  if (!fs.existsSync(artifactPath)) {
    throw new Error(
      'Missing build/UniswapV2Factory.json — run `yarn compile` in the repo root first.'
    )
  }
  const raw = fs.readFileSync(artifactPath, 'utf8')
  return JSON.parse(raw)
}

function getBytecode(artifact) {
  if (artifact.bytecode && typeof artifact.bytecode === 'string') {
    return artifact.bytecode.startsWith('0x') ? artifact.bytecode : '0x' + artifact.bytecode
  }
  const obj = artifact.evm && artifact.evm.bytecode && artifact.evm.bytecode.object
  if (obj) {
    return obj.startsWith('0x') ? obj : '0x' + obj
  }
  throw new Error('Could not find bytecode in UniswapV2Factory.json')
}

async function main() {
  const rpcUrl = process.env.RPC_URL
  if (!rpcUrl) {
    throw new Error('Set RPC_URL in .env (e.g. Sepolia HTTPS endpoint)')
  }

  const provider = new ethers.providers.JsonRpcProvider(rpcUrl)
  const network = await provider.getNetwork()
  console.log('Network chainId:', network.chainId.toString())

  const wallet = new ethers.Wallet(normalizePrivateKey(process.env.PRIVATE_KEY), provider)
  console.log('Deployer:', wallet.address)

  const balance = await wallet.getBalance()
  if (balance.isZero()) {
    console.warn('Warning: deployer balance is 0 — fund this address with test ETH for gas.')
  }

  const artifact = loadArtifact()
  const bytecode = getBytecode(artifact)
  const factory = new ethers.ContractFactory(artifact.abi, bytecode, wallet)

  const feeToSetter =
    process.env.FEE_TO_SETTER && ethers.utils.isAddress(process.env.FEE_TO_SETTER)
      ? process.env.FEE_TO_SETTER
      : wallet.address
  console.log('feeToSetter:', feeToSetter)

  console.log('Sending deployment transaction...')
  const contract = await factory.deploy(feeToSetter)
  console.log('Tx hash:', contract.deployTransaction.hash)
  await contract.deployed()
  console.log('UniswapV2Factory deployed to:', contract.address)
}

main().catch((err) => {
  console.error(err.message || err)
  process.exit(1)
})
