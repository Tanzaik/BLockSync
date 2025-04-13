# BLockSync
TLDR; a simple cloud storage system that syncs files.

BlockSync is an efficient cloud file sync system built in Golang that syncs only the modified parts of files by using block-level deduplication. It is designed to reduce bandwidth usage, optimize performance, and scale to large datasets — making it ideal for developers, system administrators, and distributed systems.

---

## 🚀 Features

- 🔹 **CLI-based file syncing**
- 🧱 **Fixed-size block splitting** (4KB per block)
- 🔐 **Content-based hashing** with SHA-256
- ☁️ **AWS S3 integration** for cloud storage
- 📦 **Manifest file tracking** to avoid redundant uploads
- 🧠 **Block-level deduplication** to sync only modified segments
- 🧰 **Modular codebase** ready for future extensions like RPC control and file restoration

---

## 📁 Project Structure

## 🔧 How It Works

1. The file is divided into **4KB blocks**
2. Each block is **hashed using SHA-256**
3. A **manifest file** is downloaded from S3 (if it exists)
4. Only **new or changed blocks** are uploaded to `s3://<bucket>/blocks/<blockhash>`
5. A new manifest is generated and stored in `s3://<bucket>/manifests/<filename>.json`

This ensures that only the necessary data is transmitted, dramatically reducing sync time and bandwidth usage.

---



Usage
1. Install dependencies
bash
Copy
Edit
go mod tidy
2. Run the CLI
bash
Copy
Edit
go run main.go <s3-bucket-name> <local-file-path>
Example:
bash
Copy
Edit
go run main.go blocksync-bucket resume.pdf


Why Block-Level Deduplication?
Traditional sync tools re-upload entire files even when only a few bytes change. BlockSync avoids that by:

Detecting changes at the block level

Reusing identical blocks already stored in the cloud

Improving sync performance and cloud storage efficiency

Enabling future support for versioning, rollback, and incremental backups

🔮 Roadmap (Future Phases)
🖥️ RPC server for remote control (using Go’s net/rpc)

♻️ Resumable uploads with partial block caching

⬇️ Reverse sync / file reconstruction from S3 blocks + manifest

📈 Content-defined chunking for smarter block detection (like rsync)

🧑‍💻 Web dashboard or GUI for visual file monitoring

🌐 AWS Integration
BlockSync uses the AWS SDK for Go and stores data in the following layout:

s3://<bucket>/blocks/<sha256> — for raw blocks

s3://<bucket>/manifests/<filename>.json — for the manifest file

You must configure your AWS credentials and region via environment variables or shared AWS config files.

🧠 Tech Stack
Language: Golang

Cloud: AWS S3

Hashing: SHA-256

Sync Strategy: Block-level deduplication

Transport: CLI-based; RPC planned

## 🏁 Conclusion

BlockSync represents a foundational step toward building smarter, faster, and more scalable file synchronization systems. By combining Golang's performance with AWS S3’s reliability and an intelligent block-level deduplication strategy, this project minimizes redundancy and maximizes efficiency. Phase 1 demonstrates how simple design principles — content-addressable storage, hashing, and manifest tracking — can enable high-impact optimizations in data transfer and storage. BlockSync is designed with extensibility in mind, setting the stage for future enhancements such as remote RPC control, resumable uploads, file reconstruction, content-defined chunking, and full versioning support. Whether you're syncing gigabytes of logs or collaborating on evolving datasets, BlockSync offers a fast, efficient, and flexible foundation to build on.

