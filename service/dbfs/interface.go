// Code generated from OpenAPI specs by Databricks SDK Generator. DO NOT EDIT.

package dbfs

import (
	"context"
)

// DBFS API makes it simple to interact with various data sources without having
// to include a users credentials every time to read a file.
type DbfsService interface {

	// Append data block
	//
	// Appends a block of data to the stream specified by the input handle. If
	// the handle does not exist, this call will throw an exception with
	// ``RESOURCE_DOES_NOT_EXIST``.
	//
	// If the block of data exceeds 1 MB, this call will throw an exception with
	// ``MAX_BLOCK_SIZE_EXCEEDED``.
	//
	// Example of request: ``` { "data": "ZGF0YWJyaWNrcwo=", "handle": 7904256 }
	// ```
	AddBlock(ctx context.Context, request AddBlock) error

	// Close the stream
	//
	// Closes the stream specified by the input handle. If the handle does not
	// exist, this call throws an exception with ``RESOURCE_DOES_NOT_EXIST``.
	Close(ctx context.Context, request Close) error

	// Open a stream
	//
	// "Opens a stream to write to a file and returns a handle to this stream.
	// There is a 10 minute idle timeout on this handle. If a file or directory
	// already exists on the given path and __overwrite__ is set to `false`,
	// this call throws an exception with ``RESOURCE_ALREADY_EXISTS``.
	//
	// A typical workflow for file upload would be:
	//
	// 1. Issue a `create` call and get a handle. 2. Issue one or more
	// `add-block` calls with the handle you have. 3. Issue a `close` call with
	// the handle you have.
	Create(ctx context.Context, request Create) (*CreateResponse, error)

	// Delete a file/directory
	//
	// "Deletes the file or directory (optionally, recursively delete all files
	// in the directory).
	//
	// This all throws an exception with ``IO_ERROR`` if the path is a non-empty
	// directory and recursive is set to `false` or other similar errors.",
	Delete(ctx context.Context, request Delete) error

	// Get the information of a file or directory
	//
	// Gets the file information for a file or directory. If the file or
	// directory does not exist, this call throws an exception with
	// ``RESOURCE_DOES_NOT_EXIST``.
	GetStatus(ctx context.Context, request GetStatusRequest) (*FileInfo, error)

	// List directory contents or file details
	//
	// Lists the contents of a directory, or details of a file. If the file or
	// directory does not exist, this call throws an exception with
	// ``RESOURCE_DOES_NOT_EXIST``.
	//
	// Example of reply:
	//
	// ``` { "files": [ { "path": "/a.cpp", "is_dir": false, "file_size\": 261
	// }, { "path": "/databricks-results", "is_dir": true, "file_size\": 0 } ] }
	// ```
	//
	// Use ListAll() to get all FileInfo instances
	List(ctx context.Context, request ListRequest) (*ListStatusResponse, error)

	// Create a directory
	//
	// Creates the given directory and necessary parent directories if they do
	// not exist. If a file (not a directory) exists at any prefix of the input
	// path, this call throws an exception with ``RESOURCE_ALREADY_EXISTS``.
	// **Note**: If this operation fails, it might have succeeded in creating
	// some of the necessary parent directories.",
	Mkdirs(ctx context.Context, request MkDirs) error

	// Move a file
	//
	// Moves a file from one location to another location within DBFS. If the
	// source file does not exist, this call throws an exception with
	// ``RESOURCE_DOES_NOT_EXIST``. If a file already exists in the destination
	// path, this call throws an exception with ``RESOURCE_ALREADY_EXISTS``. If
	// the given source path is a directory, this call always recursively moves
	// all files.",
	Move(ctx context.Context, request Move) error

	// Upload a file
	//
	// Uploads a file through the use of multipart form post. It is mainly used
	// for streaming uploads, but can also be used as a convenient single call
	// for data upload.
	//
	// Example usage:
	//
	// ``` curl -u USER:PASS -F contents=@localsrc -F
	// path="https://XX.cloud.databricks.com/api/2.0/dbfs/put" ```
	//
	// Please note that ``localsrc`` is the path to a local file to upload and
	// this usage is only supported with multipart form post (such as using -F
	// or --form with curl).
	//
	// Alternatively you can pass contents as base64 string.
	//
	// The amount of data that can be passed (when not streaming) using the
	// __contents__ parameter is limited to 1 MB. ``MAX_BLOCK_SIZE_EXCEEDED``
	// will be thrown if this limit is exceeded.
	//
	// If you want to upload large files, use the streaming upload. For details,
	// see :method:create, :method:addBlock, :method:close.
	Put(ctx context.Context, request Put) error

	// Get the contents of a file
	//
	// "Returns the contents of a file. If the file does not exist, this call
	// throws an exception with ``RESOURCE_DOES_NOT_EXIST``. If the path is a
	// directory, the read length is negative, or if the offset is negative,
	// this call throws an exception with ``INVALID_PARAMETER_VALUE``. If the
	// read length exceeds 1 MB, this call throws an\nexception with
	// ``MAX_READ_SIZE_EXCEEDED``.
	//
	// If ``offset + length`` exceeds the number of bytes in a file, it reads
	// the contents until the end of file.",
	Read(ctx context.Context, request ReadRequest) (*ReadResponse, error)
}
