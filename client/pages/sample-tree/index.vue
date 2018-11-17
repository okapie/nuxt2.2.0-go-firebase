<template>
  <div class="sample-tree">
    <section class="container">
      <div>
        <h2>Tree Sample</h2>
      </div>
      <dl class="list-wrapper">
        <dt class="term-header">Item</dt>
        <dd class="description-header">Content</dd>
        <dt class="term-body">DOM tree</dt>
        <dd class="description-body">
          <div>
            <button @click="addElement">Add</button>
          </div>
          <div>
            <div v-if="$store.state.inputs.length > 0">
              <ul v-for="input in $store.state.inputs" :key="input.index">
                <li>
                  <label>TEST</label>
                  <input
                    :name="'additional_' + input.index"
                    :value="input.value"
                    placeholder="Name"
                    type="text"
                    @change="changeValue"
                  />
                  <input class="delete-button" type="button" value="delete" @click="deleteElement(input.index)" />
                </li>
              </ul>
            </div>
          </div>
        </dd>
      </dl>
      <nuxt-link to="/">Top page</nuxt-link>
    </section>
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
        counter: 0,
        inputs: this.$store.state.inputs || []
      }
    },
    methods: {
      addElement() {
        let counter = this.counter++
        this.inputs.push({
          index: this.inputs.length > 0 ? counter : 0,
          id: this.inputs.length > 0 ? counter : 0,
          value: "",
        })
        this.$store.commit("onSaveInputsValue", this.inputs)
        this._watcher.update()
      },
      changeValue(e) {
        const index = +e.target.name.split("_")[1]
        this.inputs[index].value = e.currentTarget.value
        this.$store.commit("onSaveInputsValue", this.inputs)
        this._watcher.update()
      },
      deleteElement(index) {
        this.inputs = this.inputs.filter(x => x.index !== index)
        Object.entries(this.inputs).map(([key, value], elementIndex) => {
          value.index = elementIndex
        })
        this.$store.commit("onSaveInputsValue", this.inputs)
        this._watcher.update()
      }
    },
  }
</script>
