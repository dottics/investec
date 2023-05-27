## Changelog
All notable changes to this project will be documented in this file.

The format is based on [Keep a changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]
## [0.1.0] - 2023-05-27
### Added
- The `Credentials` struct to group the credentials information together.
- Updated the `GetAccouts` method to return a pointer value.
- Updated the `GetTransactions` method to return a pointer value.

## [0.0.3] - 2023-02-11
### Added
- The `GetTransactions` method which returns all the transactions for an
  account with optional query parameters.

## [0.0.2] - 2023-02-11
### Added 
- The `Auth` method which returns the user's access token.
- The `GetAccounts` method which returns the user's accounts.

## [0.0.1] - 2023-02-11
### Added
- The `Service` type which defines the integration struct to access the methods
  to interface with Investec.

## [0.0.0] - 2023-02-11
### Added
- Initial commit to create a basic Investec Go package.
- Base64 encoding of the Basic authentication token.
