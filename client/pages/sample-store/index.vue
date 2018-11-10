<template>
  <section class="container">
    <div>
      <h2>Vuex Store Sample</h2>
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
  export default {
    methods: {
      countNumber(){
        this.$store.commit("increment")
      },
      displayText(){
        this.$store.commit("display")
      }
    },
  }
</script>

