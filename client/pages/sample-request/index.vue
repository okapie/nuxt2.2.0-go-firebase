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
    <nuxt-link to="/">Top page</nuxt-link>
  </section>
</template>

<script>
  import axios from "axios"

  export default {
    data () {
      return {
        cardList: []
      }
    },
    methods: {
      async getCardList() {
        return await axios.get("http://localhost:8000/api/v1/cards")
          .then(res => {
            this.cardList = res.data
          })
      }
    },
  }
</script>
