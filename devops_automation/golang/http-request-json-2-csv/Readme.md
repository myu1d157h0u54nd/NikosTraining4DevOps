# http-request-json-2-csv

## GOAL

Learn some basics about the architecture and components of a blockchain.

Start practices with GO (golang) for automation in devops scripting.

## Documentation investigated for study

* Blockchains
  * [What are Blockchain Nodes?](https://dydx.exchange/crypto-learning/what-are-blockchain-nodes)
    * Full nodes
    * Staking nodes
  * [ABCI Protocol](https://cosmos-network.gitbooks.io/cosmos-academy/content/cosmos-for-developers/tendermint/abci-protocol.html)
* APIs:
  * [What’s the Difference Between RPC and REST?](https://aws.amazon.com/compare/the-difference-between-rpc-and-rest/?nc1=h_ls)
* Develop environment
  * GO:
    * Install:
      * https://go.dev/doc/install
    * Using Third-Party Packages
      * https://www.digitalocean.com/community/tutorials/importing-packages-in-go#step-2-using-third-party-packages
    * Develop:
      * Functions
        * https://go.dev/tour/basics/4
      * Functions, Named return values
        * https://go.dev/tour/basics/7
      * Data:
        * JSON:
          * A Complete Guide to JSON in Golang (With Examples)
            * https://www.sohamkamani.com/golang/json/
          * Data structure
            * https://www.golangprograms.com/go-language/struct.html
      * HTTP request
        * How To Make HTTP Requests in Go
          * https://www.digitalocean.com/community/tutorials/how-to-make-http-requests-in-go
        * Making HTTP Requests in Go
          * https://medium.com/@nurettinabaci/making-http-requests-in-go-1d599569d647
        * Replacing Golang’s ioutil functions and methods
          * https://medium.com/@ButcherWithSmile/replacing-golangs-ioutil-functions-and-methods-ec9dca2c7e59
      * A Comprehensive Guide to Using JSON in Go
        * https://betterstack.com/community/guides/scaling-go/json-in-go/

## ACTION PLAN

* Look for endpoint:

  * https://docs.dydx.exchange/network/resources
* Resources for study the endpoint:

  * https://docs.cosmos.network/main/build/modules/staking#rest
  * https://docs.cosmos.network/api
* Test the API, and POCs with CURL
* Write golang script:

  * Make http request to api json
  * Process json data
  * Concatenate data from 2 APIs
    * Moniker + VotingPower + Key
  * Structure data from json to csv
  * Save results in csv

Full node endpoints:

| endpoints | Team         | URI                                          | Info         |
| --------- | ------------ | -------------------------------------------- | ------------ |
| rpc       | CosmosSpaces | `https://rpc-dydx.cosmos-spaces.cloud:443` | moniker      |
| api       | CosmosSpaces | `https://api-dydx.cosmos-spaces.cloud:443` | voting_power |

## STEPS

##### TEST THE API AND POCS WITH CURL

curl -X GET "https://api-dydx.cosmos-spaces.cloud:443/cosmos/staking/v1beta1/validators" -H "accept: application/json" | jq

curl -X GET "https://rpc-dydx.cosmos-spaces.cloud:443/validators" -H  "accept: application/json" | jq

##### Setting up the development environment

* Curl to :
  * https://rpc-dydx.cosmos-spaces.cloud/validators?height=11473919&page=1&per_page=100

    * id validator, done
    * containing Moniker(name), done
    * Voting Power of the top 10, fail
  * https://rpc-dydx.cosmos-spaces.cloud:443

    * id validator, done
    * containing Moniker(name), fail
    * Voting Power of the top 10, done


#### Write golang script:

##### GO, Download and install

wget https://go.dev/dl/go1.22.1.linux-amd64.tar.gz

rm -rf /usr/local/go

sudo tar -C /usr/local -xvzf go1.22.1.linux-amd64.tar.gz

echo "export PATH=$PATH:/usr/local/go/bin" >> ~/.profile
source ~/.profile

go version


###### GO, writing: Apps_tech_associate_task_2_v1.go

Create a new file called main.go

    vim main.go

Compile and run Go program:

    go run main.go
