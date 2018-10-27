<template>
  <section class="container">
    <div>
      <h2>Http Request Sample</h2>
    </div>
    <div class="list-wrapper">
      <ul class="list-header">
        <li>Item</li>
        <li>Content</li>
      </ul>
      <ul class="list-body">
        <li>Http request</li>
        <li>
          <button @click="apiPublic">public</button>
          <button @click="apiPrivate">private</button>
          <p>{{ message }}</p>
        </li>
      </ul>
    </div>
    <div>
      <h2>Vuex Store Samples</h2>
    </div>
    <div class="list-wrapper">
      <ul class="list-header">
        <li>Item</li>
        <li>Content</li>
      </ul>
      <ul class="list-body">
        <li>Counter</li>
        <li>
          <button @click="countNumber">{{ $store.state.counter }}</button>
        </li>
      </ul>
      <ul class="list-body">
        <li>Display</li>
        <li>
          <button @click="displayText">Display</button>{{ $store.state.show }}
          <p v-if="$store.state.show">WWWWW</p>
        </li>
      </ul>
    </div>
    <nuxt-link to="/">Top page</nuxt-link>
  </section>
</template>

<script>
  import axios from "axios"
  import AppLogo from "~/components/AppLogo.vue"

  export default {
    components: {
      AppLogo
    },
    data () {
      return {
        message: "Foo!"
      }
    },
    methods: {
      countNumber(){
        this.$store.commit("increment")
      },
      displayText(){
        this.$store.commit("display")
      },
      apiPublic() {
        axios.get("http://localhost:8000/public")
          .then(res => {
            this.message = res.data
          })
      },
      apiPrivate() {
        axios.get("http://localhost:8000/private")
          .then(res => {
            this.message = res.data
          })
      }
    }
  }
</script>

<style>
  .container {
    display: block;
    text-align: center;
    margin-top: 20px;
    font-family: sans-serif;
  }
  .list-wrapper {
    margin: 0 4%;
    width: 92%;
  }
  ul {
    display: flex;
    flex-wrap: wrap;
    margin: 0;
    padding: 0;
  }
  ul.list-header > li {
    list-style: none;
    background: #b0b0b0;
    box-sizing: border-box;
    border-color: #000;
    border-style: solid;
    border-width: 1px 1px 1px 0;
    padding: 4px;
  }
  ul.list-body > li {
    height: 60px;
    text-align: center;
    list-style: none;
    background: #fff;
    box-sizing: border-box;
    border-color: #000;
    border-style: solid;
    border-width: 0 1px 1px 0;
    padding: 5px;
  }
  /** Item column */
  ul.list-header > li:nth-child(1),
  ul.list-body > li:nth-child(1) {
    width: 30%;
    border-left-width: 1px;
  }
  /** Content column */
  ul.list-header > li:nth-child(2),
  ul.list-body > li:nth-child(2) {
    width: 70%;
  }
  ul.list-header > li:last-child {
    border-rignt-width: 1px;
  }
  ul.list-body > li:nth-child(2n+1) {
    border-rignt-width: 1px;
  }
  p {
    color: #ff4500;
    font-size: 12px;
  }
</style>
