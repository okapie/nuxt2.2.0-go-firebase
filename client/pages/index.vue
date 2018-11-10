<template>
  <section class="container">
    <div>
      <app-logo/>
      <h1 class="title">
        nuxt2.2.0-vue
      </h1>
      <h2 class="subtitle">
        Nuxt.js project
      </h2>
      <div class="links">
        <nuxt-link 
          to="/users/" 
          class="button--green">Users</nuxt-link>
        <nuxt-link 
          to="/sample-request/"
          class="button--green">Http Request</nuxt-link>
        <nuxt-link
          to="/sample-store/"
          class="button--green">Vuex Store</nuxt-link>
      </div>
    </div>
  </section>
</template>

<script>
import AppLogo from "~/components/AppLogo.vue"

export default {
  components: {
    AppLogo
  }
}
</script>

<style>
.container {
  margin-top: 20px;
  text-align: center;
}

.title {
  font-family: "Quicksand", "Source Sans Pro", -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, "Helvetica Neue", Arial, sans-serif; /* 1 */
  display: block;
  font-weight: 300;
  font-size: 100px;
  color: #35495e;
  letter-spacing: 1px;
}

.subtitle {
  font-weight: 300;
  font-size: 42px;
  color: #526488;
  word-spacing: 5px;
  padding-bottom: 15px;
}

.links {
  padding-top: 15px;
}
</style>

