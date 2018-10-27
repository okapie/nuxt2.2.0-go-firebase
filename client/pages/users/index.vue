<template>
  <div class="container">
    <div>
      <h2>Users</h2>
    </div>
    <div class="list-wrapper">
      <ul class="list-header">
        <li>ID</li>
        <li>Name</li>
        <li>Email</li>
        <li>Detail</li>
      </ul>
      <ul v-for="user in users" :key="user.id" class="list-body">
        <li>{{ user.id }}</li>
        <li>{{ user.name }}</li>
        <li>{{ user.email }}</li>
        <li>
          <button class="detail-button">
            <nuxt-link :to="'/users/'+user.id">Detail</nuxt-link>
          </button>
        </li>
      </ul>
    </div>
    <nuxt-link to="/">Top page</nuxt-link>
  </div>
</template>

<script>
import axios from "axios"

export default {
  async asyncData() {
    const { data } = await axios.get("https://jsonplaceholder.typicode.com/users")
    return { users: data }
  }
}
</script>

<style scoped>
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
    text-align: left;
    list-style: none;
    background: #fff;
    box-sizing: border-box;
    border-color: #000;
    border-style: solid;
    border-width: 0 1px 1px 0;
    padding: 5px;
  }
  /** ID column */
  ul.list-header > li:nth-child(1),
  ul.list-body > li:nth-child(1) {
    width: 10%;
    border-left-width: 1px;
  }
  /** Name column */
  ul.list-header > li:nth-child(2),
  ul.list-body > li:nth-child(2) {
    width: 30%;
  }
  /** Email column */
  ul.list-header > li:nth-child(3),
  ul.list-body > li:nth-child(3) {
    width: 40%;
  }
  /** Detail column */
  ul.list-header > li:nth-child(4),
  ul.list-body > li:nth-child(4) {
    width: 20%;
    text-align: center;
  }
  ul.list-header > li:last-child {
    border-rignt-width: 1px;
  }
  ul.list-body > li:nth-child(2n+1) {
    border-rignt-width: 1px;
  }
  .detail-button {
    width: 90%;
    height: 24px;
  }
</style>
