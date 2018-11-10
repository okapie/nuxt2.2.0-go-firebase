<template>
  <div class="sample-tree">
    <section class="container">
      <div>
        <h2>Tree Sample</h2>
      </div>
      <dl class="list-wrapper">
        <dt class="term-header">Item</dt>
        <dd class="description-header">Content</dd>
        <dt class="term-body">Single component view</dt>
        <dd class="description-body">
          <CloneItem />
        </dd>
        <dt class="term-body">DOM tree: innerHTML</dt>
        <dd class="description-body">
          <div>
            <button @click="addElement">Add</button>
          </div>
          <div id="parent"></div>
        </dd>
      </dl>
      <nuxt-link to="/">Top page</nuxt-link>
    </section>
    <script>
      (() => {
      window.parent = document.getElementById("parent")
      })();
    </script>
  </div>
</template>

<script>
  import CloneItem from "~/components/CloneItem.vue"

  export default {
    components: {
      CloneItem
    },
    methods: {
      addElement() {
        if (process.browser) {
          // FIXME: Can not insert component with innerHTML.
          // window.parent.innerHTML += "<CloneItem />";
          window.parent.innerHTML +=
            "<div class=\"clone-item\">\n" +
            "  <div class=\"main\">\n" +
            "    <label>TEST</label>\n" +
            "    <input type=\"text\" placeholder=\"Name\" />\n" +
            "  </div>\n" +
            "</div>";
        }
      }
    },
  }
</script>
