## About Groovy Layer-1 Blockchain: A Blockchain-Powered Journey to the 70s

Groovy is more than just a website; it's a fusion of the past and the future. We're a platform celebrating the vibrant 1970s while leveraging blockchain technology to revolutionize digital experiences and the cannabis industry.

**🚀Our Mission:🚀**

To blend the nostalgia of the 70s with cutting-edge blockchain technology to create unique, engaging, and secure experiences for our users.

**🤝What We Offer:🤝**

* **Groovy Desktop Web App:** Immerse yourself in the 70s through curated content, interactive timelines, and playlists, all wrapped in a retro-inspired design.
* **NFT-QR Tag Authentication:** Explore how we're enhancing the cannabis industry with blockchain-based NFT-QR tags for transparent product authentication and tracking.
* **Groovy Blockchain (Built on Cosmos):** Experience the power of our custom blockchain, built using the Cosmos SDK and Tendermint consensus engine. We leverage Ignite CLI for efficient development and deployment.
* **Groovy Blockchain Explorer:**  Delve into the inner workings of our blockchain at [https://eyes.groovy.click](https://eyes.groovy.click), where you can view transactions, blocks, and other blockchain data.
* **P2P Geno-Library:** Discover how our decentralized peer-to-peer network fosters community building and secure data sharing within the Groovy ecosystem.
* **Rewards & Gamification:** Get involved in our community and earn rewards through our innovative gamification programs.

**🛠️Our Technology:🛠️**

* **Blockchain Infrastructure:**
    * Cosmos SDK: A modular framework for building custom blockchains.
    * Tendermint:  A Byzantine fault-tolerant consensus engine for secure and reliable blockchain operation.
    * Ignite CLI: A powerful command-line interface for simplifying blockchain development and deployment.
* **Web Technologies:**
    * React: A popular JavaScript library for building user interfaces.
    * HTML5, CSS3, JavaScript: The foundation of our web app development.
* **P2P Network:**
    * Geno-Library: Our custom-built P2P software for secure data sharing and community building.

**Join the Groovy Movement:**

Whether you're a 70s enthusiast, a blockchain aficionado, or someone interested in the future of the cannabis industry, we invite you to join our vibrant community. Explore our website, participate in our blockchain ecosystem, and experience the groovy fusion of past and future!

**☮️Keep on Groovin'!☮️** 

## Get started

```
ignite chain serve
```

`serve` command installs dependencies, builds, initializes, and starts your blockchain in development.

### Configure

Your blockchain in development can be configured with `config.yml`. To learn more, see the [Ignite CLI docs](https://docs.ignite.com).

### Web Frontend

Ignite CLI has scaffolded a Vue.js-based web app in the `vue` directory. Run the following commands to install dependencies and start the app:

```
cd vue
npm install
npm run serve
```

The frontend app is built using the `@starport/vue` and `@starport/vuex` packages. For details, see the [monorepo for Ignite front-end development](https://github.com/ignite/web).

## Release
To release a new version of your blockchain, create and push a new tag with `v` prefix. A new draft release with the configured targets will be created.

```
git tag v0.1
git push origin v0.1
```

After a draft release is created, make your final changes from the release page and publish it.

### Install
To install the latest version of your blockchain node's binary, execute the following command on your machine:

```
curl https://get.ignite.com/username/groovy@latest! | sudo bash
```
`username/groovy` should match the `username` and `repo_name` of the Github repository to which the source code was pushed. Learn more about [the install process](https://github.com/allinbits/starport-installer).

## Learn more

- [Ignite CLI](https://ignite.com/cli)
- [Tutorials](https://docs.ignite.com/guide)
- [Ignite CLI docs](https://docs.ignite.com)
- [Cosmos SDK docs](https://docs.cosmos.network)
- [Developer Chat](https://discord.gg/ignite)
