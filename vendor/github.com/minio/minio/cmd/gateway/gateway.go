/*
 * MinIO Object Storage (c) 2021 MinIO, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package gateway

import (
	// Import all gateways please keep the order

	// NAS
	// _ "github.com/minio/minio/cmd/gateway/nas"

	// Azure
	_ "github.com/minio/minio/cmd/gateway/azure"
	// S3
	// _ "github.com/minio/minio/cmd/gateway/s3"
	// gateway functionality is frozen, no new gateways are being implemented
	// or considered for upstream inclusion at this point in time. if needed
	// please keep a fork of the project.
)

// go get -u github.com/Azure/azure-pipeline-go/pipeline
// go get -u github.com/Azure/azure-storage-blob-go/azblob
// go get -u github.com/dustin/go-humanize
// go get -u github.com/minio/cli
// go get -u github.com/minio/madmin-go
// go get -u github.com/minio/minio-go/v7/pkg/policy
// go get -u github.com/minio/minio/cmd
// go get -u github.com/minio/minio/internal/logger
// go get -u github.com/minio/pkg/bucket/policy
// go get -u github.com/minio/pkg/bucket/policy/condition
// go get -u github.com/minio/pkg/env
