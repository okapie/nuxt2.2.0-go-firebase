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
      window.parentNode = window.parent.parentNode
      window.deleteElement = target => {window.parent.parentNode.removeChild(document.getElementById(target))}
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
    data () {
      return {
        counter: 0
      }
    },
    methods: {
      deleteElement(e) {
        console.log(e.currentTarget.value)
      },
      addElement() {
        if (process.browser) {
          let helper = document.createElement("div")
          const counter = this.counter++
          helper.innerHTML +=
            "<div id=\"clonedForm_" + counter + "\" class=\"clone-item\">\n" +
            "  <div class=\"main\">\n" +
            "    <label>TEST</label>\n" +
            "    <input type=\"text\" placeholder=\"Name\" value=\"" + counter + "\" />\n" +
            "    <input class=\"delete-button\" type='button' value='delete' onClick='window.deleteElement(\"clonedForm_" + counter + "\")' />" +
            "  </div>\n" +
            "</div>";

          while (helper.firstChild) {
            window.parentNode.insertBefore(helper.firstChild, window.parent)
          }
        }
      }
    },
  }
</script>
