// Code generated by "stringer -type=APIErrorCode -trimprefix=Err api-errors.go"; DO NOT EDIT.

package cmd

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[ErrNone-0]
	_ = x[ErrAccessDenied-1]
	_ = x[ErrBadDigest-2]
	_ = x[ErrEntityTooSmall-3]
	_ = x[ErrEntityTooLarge-4]
	_ = x[ErrPolicyTooLarge-5]
	_ = x[ErrIncompleteBody-6]
	_ = x[ErrInternalError-7]
	_ = x[ErrInvalidAccessKeyID-8]
	_ = x[ErrAccessKeyDisabled-9]
	_ = x[ErrInvalidArgument-10]
	_ = x[ErrInvalidBucketName-11]
	_ = x[ErrInvalidDigest-12]
	_ = x[ErrInvalidRange-13]
	_ = x[ErrInvalidRangePartNumber-14]
	_ = x[ErrInvalidCopyPartRange-15]
	_ = x[ErrInvalidCopyPartRangeSource-16]
	_ = x[ErrInvalidMaxKeys-17]
	_ = x[ErrInvalidEncodingMethod-18]
	_ = x[ErrInvalidMaxUploads-19]
	_ = x[ErrInvalidMaxParts-20]
	_ = x[ErrInvalidPartNumberMarker-21]
	_ = x[ErrInvalidPartNumber-22]
	_ = x[ErrInvalidRequestBody-23]
	_ = x[ErrInvalidCopySource-24]
	_ = x[ErrInvalidMetadataDirective-25]
	_ = x[ErrInvalidCopyDest-26]
	_ = x[ErrInvalidPolicyDocument-27]
	_ = x[ErrInvalidObjectState-28]
	_ = x[ErrMalformedXML-29]
	_ = x[ErrMissingContentLength-30]
	_ = x[ErrMissingContentMD5-31]
	_ = x[ErrMissingRequestBodyError-32]
	_ = x[ErrMissingSecurityHeader-33]
	_ = x[ErrNoSuchBucket-34]
	_ = x[ErrNoSuchBucketPolicy-35]
	_ = x[ErrNoSuchBucketLifecycle-36]
	_ = x[ErrNoSuchLifecycleConfiguration-37]
	_ = x[ErrInvalidLifecycleWithObjectLock-38]
	_ = x[ErrNoSuchBucketSSEConfig-39]
	_ = x[ErrNoSuchCORSConfiguration-40]
	_ = x[ErrNoSuchWebsiteConfiguration-41]
	_ = x[ErrReplicationConfigurationNotFoundError-42]
	_ = x[ErrRemoteDestinationNotFoundError-43]
	_ = x[ErrReplicationDestinationMissingLock-44]
	_ = x[ErrRemoteTargetNotFoundError-45]
	_ = x[ErrReplicationRemoteConnectionError-46]
	_ = x[ErrReplicationBandwidthLimitError-47]
	_ = x[ErrBucketRemoteIdenticalToSource-48]
	_ = x[ErrBucketRemoteAlreadyExists-49]
	_ = x[ErrBucketRemoteLabelInUse-50]
	_ = x[ErrBucketRemoteArnTypeInvalid-51]
	_ = x[ErrBucketRemoteArnInvalid-52]
	_ = x[ErrBucketRemoteRemoveDisallowed-53]
	_ = x[ErrRemoteTargetNotVersionedError-54]
	_ = x[ErrReplicationSourceNotVersionedError-55]
	_ = x[ErrReplicationNeedsVersioningError-56]
	_ = x[ErrReplicationBucketNeedsVersioningError-57]
	_ = x[ErrReplicationDenyEditError-58]
	_ = x[ErrRemoteTargetDenyAddError-59]
	_ = x[ErrReplicationNoExistingObjects-60]
	_ = x[ErrReplicationValidationError-61]
	_ = x[ErrReplicationPermissionCheckError-62]
	_ = x[ErrObjectRestoreAlreadyInProgress-63]
	_ = x[ErrNoSuchKey-64]
	_ = x[ErrNoSuchUpload-65]
	_ = x[ErrInvalidVersionID-66]
	_ = x[ErrNoSuchVersion-67]
	_ = x[ErrNotImplemented-68]
	_ = x[ErrPreconditionFailed-69]
	_ = x[ErrRequestTimeTooSkewed-70]
	_ = x[ErrSignatureDoesNotMatch-71]
	_ = x[ErrMethodNotAllowed-72]
	_ = x[ErrInvalidPart-73]
	_ = x[ErrInvalidPartOrder-74]
	_ = x[ErrMissingPart-75]
	_ = x[ErrAuthorizationHeaderMalformed-76]
	_ = x[ErrMalformedPOSTRequest-77]
	_ = x[ErrPOSTFileRequired-78]
	_ = x[ErrSignatureVersionNotSupported-79]
	_ = x[ErrBucketNotEmpty-80]
	_ = x[ErrAllAccessDisabled-81]
	_ = x[ErrPolicyInvalidVersion-82]
	_ = x[ErrMissingFields-83]
	_ = x[ErrMissingCredTag-84]
	_ = x[ErrCredMalformed-85]
	_ = x[ErrInvalidRegion-86]
	_ = x[ErrInvalidServiceS3-87]
	_ = x[ErrInvalidServiceSTS-88]
	_ = x[ErrInvalidRequestVersion-89]
	_ = x[ErrMissingSignTag-90]
	_ = x[ErrMissingSignHeadersTag-91]
	_ = x[ErrMalformedDate-92]
	_ = x[ErrMalformedPresignedDate-93]
	_ = x[ErrMalformedCredentialDate-94]
	_ = x[ErrMalformedExpires-95]
	_ = x[ErrNegativeExpires-96]
	_ = x[ErrAuthHeaderEmpty-97]
	_ = x[ErrExpiredPresignRequest-98]
	_ = x[ErrRequestNotReadyYet-99]
	_ = x[ErrUnsignedHeaders-100]
	_ = x[ErrMissingDateHeader-101]
	_ = x[ErrInvalidQuerySignatureAlgo-102]
	_ = x[ErrInvalidQueryParams-103]
	_ = x[ErrBucketAlreadyOwnedByYou-104]
	_ = x[ErrInvalidDuration-105]
	_ = x[ErrBucketAlreadyExists-106]
	_ = x[ErrMetadataTooLarge-107]
	_ = x[ErrUnsupportedMetadata-108]
	_ = x[ErrUnsupportedHostHeader-109]
	_ = x[ErrMaximumExpires-110]
	_ = x[ErrSlowDownRead-111]
	_ = x[ErrSlowDownWrite-112]
	_ = x[ErrMaxVersionsExceeded-113]
	_ = x[ErrInvalidPrefixMarker-114]
	_ = x[ErrBadRequest-115]
	_ = x[ErrKeyTooLongError-116]
	_ = x[ErrInvalidBucketObjectLockConfiguration-117]
	_ = x[ErrObjectLockConfigurationNotFound-118]
	_ = x[ErrObjectLockConfigurationNotAllowed-119]
	_ = x[ErrNoSuchObjectLockConfiguration-120]
	_ = x[ErrObjectLocked-121]
	_ = x[ErrInvalidRetentionDate-122]
	_ = x[ErrPastObjectLockRetainDate-123]
	_ = x[ErrUnknownWORMModeDirective-124]
	_ = x[ErrBucketTaggingNotFound-125]
	_ = x[ErrObjectLockInvalidHeaders-126]
	_ = x[ErrInvalidTagDirective-127]
	_ = x[ErrPolicyAlreadyAttached-128]
	_ = x[ErrPolicyNotAttached-129]
	_ = x[ErrExcessData-130]
	_ = x[ErrInvalidEncryptionMethod-131]
	_ = x[ErrInvalidEncryptionKeyID-132]
	_ = x[ErrInsecureSSECustomerRequest-133]
	_ = x[ErrSSEMultipartEncrypted-134]
	_ = x[ErrSSEEncryptedObject-135]
	_ = x[ErrInvalidEncryptionParameters-136]
	_ = x[ErrInvalidEncryptionParametersSSEC-137]
	_ = x[ErrInvalidSSECustomerAlgorithm-138]
	_ = x[ErrInvalidSSECustomerKey-139]
	_ = x[ErrMissingSSECustomerKey-140]
	_ = x[ErrMissingSSECustomerKeyMD5-141]
	_ = x[ErrSSECustomerKeyMD5Mismatch-142]
	_ = x[ErrInvalidSSECustomerParameters-143]
	_ = x[ErrIncompatibleEncryptionMethod-144]
	_ = x[ErrKMSNotConfigured-145]
	_ = x[ErrKMSKeyNotFoundException-146]
	_ = x[ErrKMSDefaultKeyAlreadyConfigured-147]
	_ = x[ErrNoAccessKey-148]
	_ = x[ErrInvalidToken-149]
	_ = x[ErrEventNotification-150]
	_ = x[ErrARNNotification-151]
	_ = x[ErrRegionNotification-152]
	_ = x[ErrOverlappingFilterNotification-153]
	_ = x[ErrFilterNameInvalid-154]
	_ = x[ErrFilterNamePrefix-155]
	_ = x[ErrFilterNameSuffix-156]
	_ = x[ErrFilterValueInvalid-157]
	_ = x[ErrOverlappingConfigs-158]
	_ = x[ErrUnsupportedNotification-159]
	_ = x[ErrContentSHA256Mismatch-160]
	_ = x[ErrContentChecksumMismatch-161]
	_ = x[ErrStorageFull-162]
	_ = x[ErrRequestBodyParse-163]
	_ = x[ErrObjectExistsAsDirectory-164]
	_ = x[ErrInvalidObjectName-165]
	_ = x[ErrInvalidObjectNamePrefixSlash-166]
	_ = x[ErrInvalidResourceName-167]
	_ = x[ErrInvalidLifecycleQueryParameter-168]
	_ = x[ErrServerNotInitialized-169]
	_ = x[ErrRequestTimedout-170]
	_ = x[ErrClientDisconnected-171]
	_ = x[ErrTooManyRequests-172]
	_ = x[ErrInvalidRequest-173]
	_ = x[ErrTransitionStorageClassNotFoundError-174]
	_ = x[ErrInvalidStorageClass-175]
	_ = x[ErrBackendDown-176]
	_ = x[ErrMalformedJSON-177]
	_ = x[ErrAdminNoSuchUser-178]
	_ = x[ErrAdminNoSuchUserLDAPWarn-179]
	_ = x[ErrAdminNoSuchGroup-180]
	_ = x[ErrAdminGroupNotEmpty-181]
	_ = x[ErrAdminGroupDisabled-182]
	_ = x[ErrAdminNoSuchJob-183]
	_ = x[ErrAdminNoSuchPolicy-184]
	_ = x[ErrAdminPolicyChangeAlreadyApplied-185]
	_ = x[ErrAdminInvalidArgument-186]
	_ = x[ErrAdminInvalidAccessKey-187]
	_ = x[ErrAdminInvalidSecretKey-188]
	_ = x[ErrAdminConfigNoQuorum-189]
	_ = x[ErrAdminConfigTooLarge-190]
	_ = x[ErrAdminConfigBadJSON-191]
	_ = x[ErrAdminNoSuchConfigTarget-192]
	_ = x[ErrAdminConfigEnvOverridden-193]
	_ = x[ErrAdminConfigDuplicateKeys-194]
	_ = x[ErrAdminConfigInvalidIDPType-195]
	_ = x[ErrAdminConfigLDAPNonDefaultConfigName-196]
	_ = x[ErrAdminConfigLDAPValidation-197]
	_ = x[ErrAdminConfigIDPCfgNameAlreadyExists-198]
	_ = x[ErrAdminConfigIDPCfgNameDoesNotExist-199]
	_ = x[ErrAdminCredentialsMismatch-200]
	_ = x[ErrInsecureClientRequest-201]
	_ = x[ErrObjectTampered-202]
	_ = x[ErrSiteReplicationInvalidRequest-203]
	_ = x[ErrSiteReplicationPeerResp-204]
	_ = x[ErrSiteReplicationBackendIssue-205]
	_ = x[ErrSiteReplicationServiceAccountError-206]
	_ = x[ErrSiteReplicationBucketConfigError-207]
	_ = x[ErrSiteReplicationBucketMetaError-208]
	_ = x[ErrSiteReplicationIAMError-209]
	_ = x[ErrSiteReplicationConfigMissing-210]
	_ = x[ErrSiteReplicationIAMConfigMismatch-211]
	_ = x[ErrAdminRebalanceAlreadyStarted-212]
	_ = x[ErrAdminRebalanceNotStarted-213]
	_ = x[ErrAdminBucketQuotaExceeded-214]
	_ = x[ErrAdminNoSuchQuotaConfiguration-215]
	_ = x[ErrHealNotImplemented-216]
	_ = x[ErrHealNoSuchProcess-217]
	_ = x[ErrHealInvalidClientToken-218]
	_ = x[ErrHealMissingBucket-219]
	_ = x[ErrHealAlreadyRunning-220]
	_ = x[ErrHealOverlappingPaths-221]
	_ = x[ErrIncorrectContinuationToken-222]
	_ = x[ErrEmptyRequestBody-223]
	_ = x[ErrUnsupportedFunction-224]
	_ = x[ErrInvalidExpressionType-225]
	_ = x[ErrBusy-226]
	_ = x[ErrUnauthorizedAccess-227]
	_ = x[ErrExpressionTooLong-228]
	_ = x[ErrIllegalSQLFunctionArgument-229]
	_ = x[ErrInvalidKeyPath-230]
	_ = x[ErrInvalidCompressionFormat-231]
	_ = x[ErrInvalidFileHeaderInfo-232]
	_ = x[ErrInvalidJSONType-233]
	_ = x[ErrInvalidQuoteFields-234]
	_ = x[ErrInvalidRequestParameter-235]
	_ = x[ErrInvalidDataType-236]
	_ = x[ErrInvalidTextEncoding-237]
	_ = x[ErrInvalidDataSource-238]
	_ = x[ErrInvalidTableAlias-239]
	_ = x[ErrMissingRequiredParameter-240]
	_ = x[ErrObjectSerializationConflict-241]
	_ = x[ErrUnsupportedSQLOperation-242]
	_ = x[ErrUnsupportedSQLStructure-243]
	_ = x[ErrUnsupportedSyntax-244]
	_ = x[ErrUnsupportedRangeHeader-245]
	_ = x[ErrLexerInvalidChar-246]
	_ = x[ErrLexerInvalidOperator-247]
	_ = x[ErrLexerInvalidLiteral-248]
	_ = x[ErrLexerInvalidIONLiteral-249]
	_ = x[ErrParseExpectedDatePart-250]
	_ = x[ErrParseExpectedKeyword-251]
	_ = x[ErrParseExpectedTokenType-252]
	_ = x[ErrParseExpected2TokenTypes-253]
	_ = x[ErrParseExpectedNumber-254]
	_ = x[ErrParseExpectedRightParenBuiltinFunctionCall-255]
	_ = x[ErrParseExpectedTypeName-256]
	_ = x[ErrParseExpectedWhenClause-257]
	_ = x[ErrParseUnsupportedToken-258]
	_ = x[ErrParseUnsupportedLiteralsGroupBy-259]
	_ = x[ErrParseExpectedMember-260]
	_ = x[ErrParseUnsupportedSelect-261]
	_ = x[ErrParseUnsupportedCase-262]
	_ = x[ErrParseUnsupportedCaseClause-263]
	_ = x[ErrParseUnsupportedAlias-264]
	_ = x[ErrParseUnsupportedSyntax-265]
	_ = x[ErrParseUnknownOperator-266]
	_ = x[ErrParseMissingIdentAfterAt-267]
	_ = x[ErrParseUnexpectedOperator-268]
	_ = x[ErrParseUnexpectedTerm-269]
	_ = x[ErrParseUnexpectedToken-270]
	_ = x[ErrParseUnexpectedKeyword-271]
	_ = x[ErrParseExpectedExpression-272]
	_ = x[ErrParseExpectedLeftParenAfterCast-273]
	_ = x[ErrParseExpectedLeftParenValueConstructor-274]
	_ = x[ErrParseExpectedLeftParenBuiltinFunctionCall-275]
	_ = x[ErrParseExpectedArgumentDelimiter-276]
	_ = x[ErrParseCastArity-277]
	_ = x[ErrParseInvalidTypeParam-278]
	_ = x[ErrParseEmptySelect-279]
	_ = x[ErrParseSelectMissingFrom-280]
	_ = x[ErrParseExpectedIdentForGroupName-281]
	_ = x[ErrParseExpectedIdentForAlias-282]
	_ = x[ErrParseUnsupportedCallWithStar-283]
	_ = x[ErrParseNonUnaryAgregateFunctionCall-284]
	_ = x[ErrParseMalformedJoin-285]
	_ = x[ErrParseExpectedIdentForAt-286]
	_ = x[ErrParseAsteriskIsNotAloneInSelectList-287]
	_ = x[ErrParseCannotMixSqbAndWildcardInSelectList-288]
	_ = x[ErrParseInvalidContextForWildcardInSelectList-289]
	_ = x[ErrIncorrectSQLFunctionArgumentType-290]
	_ = x[ErrValueParseFailure-291]
	_ = x[ErrEvaluatorInvalidArguments-292]
	_ = x[ErrIntegerOverflow-293]
	_ = x[ErrLikeInvalidInputs-294]
	_ = x[ErrCastFailed-295]
	_ = x[ErrInvalidCast-296]
	_ = x[ErrEvaluatorInvalidTimestampFormatPattern-297]
	_ = x[ErrEvaluatorInvalidTimestampFormatPatternSymbolForParsing-298]
	_ = x[ErrEvaluatorTimestampFormatPatternDuplicateFields-299]
	_ = x[ErrEvaluatorTimestampFormatPatternHourClockAmPmMismatch-300]
	_ = x[ErrEvaluatorUnterminatedTimestampFormatPatternToken-301]
	_ = x[ErrEvaluatorInvalidTimestampFormatPatternToken-302]
	_ = x[ErrEvaluatorInvalidTimestampFormatPatternSymbol-303]
	_ = x[ErrEvaluatorBindingDoesNotExist-304]
	_ = x[ErrMissingHeaders-305]
	_ = x[ErrInvalidColumnIndex-306]
	_ = x[ErrAdminConfigNotificationTargetsFailed-307]
	_ = x[ErrAdminProfilerNotEnabled-308]
	_ = x[ErrInvalidDecompressedSize-309]
	_ = x[ErrAddUserInvalidArgument-310]
	_ = x[ErrAdminResourceInvalidArgument-311]
	_ = x[ErrAdminAccountNotEligible-312]
	_ = x[ErrAccountNotEligible-313]
	_ = x[ErrAdminServiceAccountNotFound-314]
	_ = x[ErrPostPolicyConditionInvalidFormat-315]
	_ = x[ErrInvalidChecksum-316]
	_ = x[ErrLambdaARNInvalid-317]
	_ = x[ErrLambdaARNNotFound-318]
	_ = x[apiErrCodeEnd-319]
}

const _APIErrorCode_name = "NoneAccessDeniedBadDigestEntityTooSmallEntityTooLargePolicyTooLargeIncompleteBodyInternalErrorInvalidAccessKeyIDAccessKeyDisabledInvalidArgumentInvalidBucketNameInvalidDigestInvalidRangeInvalidRangePartNumberInvalidCopyPartRangeInvalidCopyPartRangeSourceInvalidMaxKeysInvalidEncodingMethodInvalidMaxUploadsInvalidMaxPartsInvalidPartNumberMarkerInvalidPartNumberInvalidRequestBodyInvalidCopySourceInvalidMetadataDirectiveInvalidCopyDestInvalidPolicyDocumentInvalidObjectStateMalformedXMLMissingContentLengthMissingContentMD5MissingRequestBodyErrorMissingSecurityHeaderNoSuchBucketNoSuchBucketPolicyNoSuchBucketLifecycleNoSuchLifecycleConfigurationInvalidLifecycleWithObjectLockNoSuchBucketSSEConfigNoSuchCORSConfigurationNoSuchWebsiteConfigurationReplicationConfigurationNotFoundErrorRemoteDestinationNotFoundErrorReplicationDestinationMissingLockRemoteTargetNotFoundErrorReplicationRemoteConnectionErrorReplicationBandwidthLimitErrorBucketRemoteIdenticalToSourceBucketRemoteAlreadyExistsBucketRemoteLabelInUseBucketRemoteArnTypeInvalidBucketRemoteArnInvalidBucketRemoteRemoveDisallowedRemoteTargetNotVersionedErrorReplicationSourceNotVersionedErrorReplicationNeedsVersioningErrorReplicationBucketNeedsVersioningErrorReplicationDenyEditErrorRemoteTargetDenyAddErrorReplicationNoExistingObjectsReplicationValidationErrorReplicationPermissionCheckErrorObjectRestoreAlreadyInProgressNoSuchKeyNoSuchUploadInvalidVersionIDNoSuchVersionNotImplementedPreconditionFailedRequestTimeTooSkewedSignatureDoesNotMatchMethodNotAllowedInvalidPartInvalidPartOrderMissingPartAuthorizationHeaderMalformedMalformedPOSTRequestPOSTFileRequiredSignatureVersionNotSupportedBucketNotEmptyAllAccessDisabledPolicyInvalidVersionMissingFieldsMissingCredTagCredMalformedInvalidRegionInvalidServiceS3InvalidServiceSTSInvalidRequestVersionMissingSignTagMissingSignHeadersTagMalformedDateMalformedPresignedDateMalformedCredentialDateMalformedExpiresNegativeExpiresAuthHeaderEmptyExpiredPresignRequestRequestNotReadyYetUnsignedHeadersMissingDateHeaderInvalidQuerySignatureAlgoInvalidQueryParamsBucketAlreadyOwnedByYouInvalidDurationBucketAlreadyExistsMetadataTooLargeUnsupportedMetadataUnsupportedHostHeaderMaximumExpiresSlowDownReadSlowDownWriteMaxVersionsExceededInvalidPrefixMarkerBadRequestKeyTooLongErrorInvalidBucketObjectLockConfigurationObjectLockConfigurationNotFoundObjectLockConfigurationNotAllowedNoSuchObjectLockConfigurationObjectLockedInvalidRetentionDatePastObjectLockRetainDateUnknownWORMModeDirectiveBucketTaggingNotFoundObjectLockInvalidHeadersInvalidTagDirectivePolicyAlreadyAttachedPolicyNotAttachedExcessDataInvalidEncryptionMethodInvalidEncryptionKeyIDInsecureSSECustomerRequestSSEMultipartEncryptedSSEEncryptedObjectInvalidEncryptionParametersInvalidEncryptionParametersSSECInvalidSSECustomerAlgorithmInvalidSSECustomerKeyMissingSSECustomerKeyMissingSSECustomerKeyMD5SSECustomerKeyMD5MismatchInvalidSSECustomerParametersIncompatibleEncryptionMethodKMSNotConfiguredKMSKeyNotFoundExceptionKMSDefaultKeyAlreadyConfiguredNoAccessKeyInvalidTokenEventNotificationARNNotificationRegionNotificationOverlappingFilterNotificationFilterNameInvalidFilterNamePrefixFilterNameSuffixFilterValueInvalidOverlappingConfigsUnsupportedNotificationContentSHA256MismatchContentChecksumMismatchStorageFullRequestBodyParseObjectExistsAsDirectoryInvalidObjectNameInvalidObjectNamePrefixSlashInvalidResourceNameInvalidLifecycleQueryParameterServerNotInitializedRequestTimedoutClientDisconnectedTooManyRequestsInvalidRequestTransitionStorageClassNotFoundErrorInvalidStorageClassBackendDownMalformedJSONAdminNoSuchUserAdminNoSuchUserLDAPWarnAdminNoSuchGroupAdminGroupNotEmptyAdminGroupDisabledAdminNoSuchJobAdminNoSuchPolicyAdminPolicyChangeAlreadyAppliedAdminInvalidArgumentAdminInvalidAccessKeyAdminInvalidSecretKeyAdminConfigNoQuorumAdminConfigTooLargeAdminConfigBadJSONAdminNoSuchConfigTargetAdminConfigEnvOverriddenAdminConfigDuplicateKeysAdminConfigInvalidIDPTypeAdminConfigLDAPNonDefaultConfigNameAdminConfigLDAPValidationAdminConfigIDPCfgNameAlreadyExistsAdminConfigIDPCfgNameDoesNotExistAdminCredentialsMismatchInsecureClientRequestObjectTamperedSiteReplicationInvalidRequestSiteReplicationPeerRespSiteReplicationBackendIssueSiteReplicationServiceAccountErrorSiteReplicationBucketConfigErrorSiteReplicationBucketMetaErrorSiteReplicationIAMErrorSiteReplicationConfigMissingSiteReplicationIAMConfigMismatchAdminRebalanceAlreadyStartedAdminRebalanceNotStartedAdminBucketQuotaExceededAdminNoSuchQuotaConfigurationHealNotImplementedHealNoSuchProcessHealInvalidClientTokenHealMissingBucketHealAlreadyRunningHealOverlappingPathsIncorrectContinuationTokenEmptyRequestBodyUnsupportedFunctionInvalidExpressionTypeBusyUnauthorizedAccessExpressionTooLongIllegalSQLFunctionArgumentInvalidKeyPathInvalidCompressionFormatInvalidFileHeaderInfoInvalidJSONTypeInvalidQuoteFieldsInvalidRequestParameterInvalidDataTypeInvalidTextEncodingInvalidDataSourceInvalidTableAliasMissingRequiredParameterObjectSerializationConflictUnsupportedSQLOperationUnsupportedSQLStructureUnsupportedSyntaxUnsupportedRangeHeaderLexerInvalidCharLexerInvalidOperatorLexerInvalidLiteralLexerInvalidIONLiteralParseExpectedDatePartParseExpectedKeywordParseExpectedTokenTypeParseExpected2TokenTypesParseExpectedNumberParseExpectedRightParenBuiltinFunctionCallParseExpectedTypeNameParseExpectedWhenClauseParseUnsupportedTokenParseUnsupportedLiteralsGroupByParseExpectedMemberParseUnsupportedSelectParseUnsupportedCaseParseUnsupportedCaseClauseParseUnsupportedAliasParseUnsupportedSyntaxParseUnknownOperatorParseMissingIdentAfterAtParseUnexpectedOperatorParseUnexpectedTermParseUnexpectedTokenParseUnexpectedKeywordParseExpectedExpressionParseExpectedLeftParenAfterCastParseExpectedLeftParenValueConstructorParseExpectedLeftParenBuiltinFunctionCallParseExpectedArgumentDelimiterParseCastArityParseInvalidTypeParamParseEmptySelectParseSelectMissingFromParseExpectedIdentForGroupNameParseExpectedIdentForAliasParseUnsupportedCallWithStarParseNonUnaryAgregateFunctionCallParseMalformedJoinParseExpectedIdentForAtParseAsteriskIsNotAloneInSelectListParseCannotMixSqbAndWildcardInSelectListParseInvalidContextForWildcardInSelectListIncorrectSQLFunctionArgumentTypeValueParseFailureEvaluatorInvalidArgumentsIntegerOverflowLikeInvalidInputsCastFailedInvalidCastEvaluatorInvalidTimestampFormatPatternEvaluatorInvalidTimestampFormatPatternSymbolForParsingEvaluatorTimestampFormatPatternDuplicateFieldsEvaluatorTimestampFormatPatternHourClockAmPmMismatchEvaluatorUnterminatedTimestampFormatPatternTokenEvaluatorInvalidTimestampFormatPatternTokenEvaluatorInvalidTimestampFormatPatternSymbolEvaluatorBindingDoesNotExistMissingHeadersInvalidColumnIndexAdminConfigNotificationTargetsFailedAdminProfilerNotEnabledInvalidDecompressedSizeAddUserInvalidArgumentAdminResourceInvalidArgumentAdminAccountNotEligibleAccountNotEligibleAdminServiceAccountNotFoundPostPolicyConditionInvalidFormatInvalidChecksumLambdaARNInvalidLambdaARNNotFoundapiErrCodeEnd"

var _APIErrorCode_index = [...]uint16{0, 4, 16, 25, 39, 53, 67, 81, 94, 112, 129, 144, 161, 174, 186, 208, 228, 254, 268, 289, 306, 321, 344, 361, 379, 396, 420, 435, 456, 474, 486, 506, 523, 546, 567, 579, 597, 618, 646, 676, 697, 720, 746, 783, 813, 846, 871, 903, 933, 962, 987, 1009, 1035, 1057, 1085, 1114, 1148, 1179, 1216, 1240, 1264, 1292, 1318, 1349, 1379, 1388, 1400, 1416, 1429, 1443, 1461, 1481, 1502, 1518, 1529, 1545, 1556, 1584, 1604, 1620, 1648, 1662, 1679, 1699, 1712, 1726, 1739, 1752, 1768, 1785, 1806, 1820, 1841, 1854, 1876, 1899, 1915, 1930, 1945, 1966, 1984, 1999, 2016, 2041, 2059, 2082, 2097, 2116, 2132, 2151, 2172, 2186, 2198, 2211, 2230, 2249, 2259, 2274, 2310, 2341, 2374, 2403, 2415, 2435, 2459, 2483, 2504, 2528, 2547, 2568, 2585, 2595, 2618, 2640, 2666, 2687, 2705, 2732, 2763, 2790, 2811, 2832, 2856, 2881, 2909, 2937, 2953, 2976, 3006, 3017, 3029, 3046, 3061, 3079, 3108, 3125, 3141, 3157, 3175, 3193, 3216, 3237, 3260, 3271, 3287, 3310, 3327, 3355, 3374, 3404, 3424, 3439, 3457, 3472, 3486, 3521, 3540, 3551, 3564, 3579, 3602, 3618, 3636, 3654, 3668, 3685, 3716, 3736, 3757, 3778, 3797, 3816, 3834, 3857, 3881, 3905, 3930, 3965, 3990, 4024, 4057, 4081, 4102, 4116, 4145, 4168, 4195, 4229, 4261, 4291, 4314, 4342, 4374, 4402, 4426, 4450, 4479, 4497, 4514, 4536, 4553, 4571, 4591, 4617, 4633, 4652, 4673, 4677, 4695, 4712, 4738, 4752, 4776, 4797, 4812, 4830, 4853, 4868, 4887, 4904, 4921, 4945, 4972, 4995, 5018, 5035, 5057, 5073, 5093, 5112, 5134, 5155, 5175, 5197, 5221, 5240, 5282, 5303, 5326, 5347, 5378, 5397, 5419, 5439, 5465, 5486, 5508, 5528, 5552, 5575, 5594, 5614, 5636, 5659, 5690, 5728, 5769, 5799, 5813, 5834, 5850, 5872, 5902, 5928, 5956, 5989, 6007, 6030, 6065, 6105, 6147, 6179, 6196, 6221, 6236, 6253, 6263, 6274, 6312, 6366, 6412, 6464, 6512, 6555, 6599, 6627, 6641, 6659, 6695, 6718, 6741, 6763, 6791, 6814, 6832, 6859, 6891, 6906, 6922, 6939, 6952}

func (i APIErrorCode) String() string {
	if i < 0 || i >= APIErrorCode(len(_APIErrorCode_index)-1) {
		return "APIErrorCode(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _APIErrorCode_name[_APIErrorCode_index[i]:_APIErrorCode_index[i+1]]
}