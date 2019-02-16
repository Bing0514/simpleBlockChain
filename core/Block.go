package core

import (
	"crypto/sha256"
	"encoding/hex"
	"time"
)

type Block struct {
	//区块头
	Index int64 // 区块编号 - 代表区块在区块链中的位置
	Timestamp int64 //区块时间戳 - 区块创建的时间
	PrevBlockHash string //上一个区块的hash值
	Hash string //当前区块的hash值

//	区块体
	Data string //区块数据
}

/**
计算hash
 */
func calculateHash(b Block) string{
	blockData := string(b.Index) + string(b.Timestamp) + b.PrevBlockHash + b.Data
	hashInButes := sha256.Sum256([]byte(blockData)) //blockData是一个字符串，注意这里要转换层字节切片
	return hex.EncodeToString(hashInButes[:])
}

/**
生成新的区块 - 数据是基于前一个区块的
 */
func GenerateNewBlock(preBlock Block, data string) Block{
	newBlock := Block{}
	newBlock.Index = preBlock.Index+1
	newBlock.PrevBlockHash = preBlock.Hash
	newBlock.Timestamp = time.Now().Unix()
	newBlock.Data = data
	newBlock.Hash = calculateHash(newBlock)
	return newBlock
}

/**
生成创世区块
index是0
hash是一个空值
 */
func GenerateGenesesBlock() Block{
	preBlock := Block{} //这个父区块是不存在的，是为了函数复用
	preBlock.Index = -1
	preBlock.Hash = ""
	return GenerateNewBlock(preBlock,"Genesis Block")
}