## Alibaba Cloud SDK Authentication Methods

According to the [Alibaba Cloud SDK documentation](https://github.com/aliyun/alibaba-cloud-sdk-go/blob/master/docs/2-Client-EN.md), there are several methods to create a client in the SDK:

1. **AccessKey Client:** [Learn more](https://github.com/aliyun/alibaba-cloud-sdk-go/blob/master/docs/2-Client-EN.md#accesskey-client)

2. **STS Client:** [Learn more](https://github.com/aliyun/alibaba-cloud-sdk-go/blob/master/docs/2-Client-EN.md#sts-client)

3. **RamRoleArn Client:** [Learn more](https://github.com/aliyun/alibaba-cloud-sdk-go/blob/master/docs/2-Client-EN.md#ramrolearn-client)

4. **EcsRamRole Client:** [Learn more](https://github.com/aliyun/alibaba-cloud-sdk-go/blob/master/docs/2-Client-EN.md#ecsramrole-client)

5. **Bearer Token Client:** [Learn more](https://github.com/aliyun/alibaba-cloud-sdk-go/blob/master/docs/2-Client-EN.md#bearer-token-client)

6. **RsaKeyPair Client:** [Learn more](https://github.com/aliyun/alibaba-cloud-sdk-go/blob/master/docs/2-Client-EN.md#rsakeypair-client)

7. **Default Credential Provider Chain:** [Learn more](https://github.com/aliyun/alibaba-cloud-sdk-go/blob/master/docs/2-Client-EN.md#default-credential-provider-chain)

### Current Authentication Implementation:

- We have only implemented Access Key (AK) authentication so far.

- Each service has a function called `NewClientWithOptions` that generates a client based on the provided credential configuration.

- The Credential Provider Chain typically searches for credentials in the following order:

  - Environment Variables

  - Configuration Files

  - ECS Role

  - Instance RAM Role

## Alibaba Cloud CLI Authentication Methods

In the Alibaba Cloud CLI, you can configure authentication methods (AK, StsToken, RamRoleArn, and EcsRamRole) in four ways as outlined in the [CLI documentation](https://github.com/aliyun/aliyun-cli?tab=readme-ov-file#configure-authentication-methods). Additionally, there is an 'External' authentication method, which presents some complexities.

## Implementation Approaches

### Option 1:

- Create a function to obtain credential configurations, similar to the AWS plugin mechanism.

- Use the `NewClientWithOptions` function to create the client for each service.

### Option 2:

- Manually determine the configuration arguments set by the user in the CLI credential path `~/.aliyun/config.json`.

- Select the authentication method based on the user's provided configuration.

- Create client creation logic based on the user-provided details or use `NewClientWithOptions` to obtain client providers based on the profile.

## Decision on Adopting the Approaches

Considering the Alibaba Cloud CLI stores credentials at `~/.aliyun/config.json`, while the SDK uses `~/.alibabacloud/credentials` in different formats ([documentation reference](https://github.com/aliyun/alibaba-cloud-sdk-go/blob/master/docs/2-Client-EN.md)), an issue was raised by users ([Issue #451](https://github.com/aliyun/alibaba-cloud-sdk-go/issues/451)) which influenced our decision against Option 1.

The SDK supports more authentication mechanisms than the CLI. The functionality of additional authentication mechanisms unsupported by the Alibaba Cloud CLI is unclear.

Based on these observations, Option 2 seems more feasible for implementing the profile-based authentication mechanism.

## Implementation Steps for Profile Authentication

1. Load the credential file from `~/.aliyun/config.json`.

2. Retrieve the profile name set in the `.spc` file.

3. Fetch the credential details for the specified profile in the `.spc` file.

4. Parse the profile credentials into a Go structure.

5. Generate the credential provider based on the authentication mode set for that profile.

6. Return the credential provider details for client creation across services.

## Ordering of authentication mechanism

1. Environment variable authentication.

2. Authenticate with given profile in .spc file

3. Authenticate with AK and SK provided in .spc file
