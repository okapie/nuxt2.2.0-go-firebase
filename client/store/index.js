import Vue from "vue"
import Vuex from "vuex"

Vue.use(Vuex)

const createStore = () => {
  return new Vuex.Store({
    state: () => ({
      counter: 0,
      formCounter: 0,
      show: false,
      inputs: []
    }),
    mutations: {
      increment (state) {
        state.counter++
      },
      display (state) {
        state.show = !state.show
        console.log(state.show)
      },
      onSaveInputsValue (state, value) {
        state.inputs = value
      }
    }
  })
}

export default createStore
