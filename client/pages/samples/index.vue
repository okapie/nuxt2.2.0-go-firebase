<template>
  <section class="container">
    <div>
      <h2>Http Request Sample</h2>
    </div>
    <dl class="list-wrapper">
      <dt class="term-header">Item</dt>
      <dd class="description-header">Content</dd>
      <dt class="term-body">Http request</dt>
      <dd class="description-body">
        <button @click="getCardList">Show Card List</button>
        <div class="list-detail">
          <div>
            <span>[ with ':src' attribute ]</span>
            <ul v-for="item in cardList" v-bind:key="item.id">
              <li>ID: {{ item.id }}</li>
              <li>NAME: {{ item.name }}</li>
              <li v-if="item.uploadedImage !==''">
                <img :src="require(`~/assets/images/${item.uploadedImage}`)">
              </li>
            </ul>
          </div>
          <div>
            <span>[ with Lazy load ]</span>
            <ul v-for="item in cardList" v-bind:key="item.id">
              <li>ID: {{ item.id }}</li>
              <li>NAME: {{ item.name }}</li>
              <li v-if="item.uploadedImage !==''">
                <img v-lazy="require(`~/assets/images/${item.uploadedImage}`)">
              </li>
            </ul>
          </div>
        </div>
      </dd>
    </dl>
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
        cardList: []
      }
    },
    methods: {
      countNumber(){
        this.$store.commit("increment")
      },
      displayText(){
        this.$store.commit("display")
      },
      async getCardList() {
        return await axios.get("http://localhost:8000/api/v1/cards")
          .then(res => {
            this.cardList = res.data
          })
      }
    },
  }
</script>

<style>
  .container {
    display: block;
    text-align: center;
    margin-top: 20px;
    font-family: sans-serif;
  }
  dl {
    display:flex;
    flex-wrap: wrap;
    border-top: none;
  }
  dt {
    background: #ddd;
    width: 30%;
    padding: 10px;
    box-sizing: border-box;
    border-top: 1px solid #ccc;
  }
  dd {
    padding: 10px;
    margin: 0;
    border-left: 1px solid #ccc;
    border-top: 1px solid #ccc;
    width: 70%;
    background: #fff;
    box-sizing: border-box;
  }
  dt.term-body,
  dd.description-body {
    height: 420px;
    text-align: center;
    background: #fff;
    box-sizing: border-box;
    border-color: #000;
    border-style: solid;
    border-width: 0 1px 1px 0;
    padding: 5px;
  }
  .list-wrapper {
    margin: 0 4%;
    width: 92%;
  }
  ul {
    flex-wrap: wrap;
    margin: 0;
    padding: 0;
  }
  li {
    list-style: none;
  }
  ul.list-header,
  ul.list-body {
    display: flex;
  }
  ul.list-header > li,
  dt.term-header,
  dd.description-header {
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
    background: #fff;
    box-sizing: border-box;
    border-color: #000;
    border-style: solid;
    border-width: 0 1px 1px 0;
    padding: 5px;
  }
  /** Item column */
  ul.list-header > li:nth-child(1),
  ul.list-body > li:nth-child(1),
  dt.term-header,
  dt.term-body {
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
  .list-detail {
    display: flex;
    height: 92%;
  }
  .list-detail > div {
    text-align: center;
    width: 50%;
    border-top: 1px #000 solid;
  }
  .list-detail > div:nth-child(1) {
    border-right: 1px #000 solid;
  }
  p {
    color: #ff4500;
    font-size: 12px;
  }
</style>
