import Vue from "vue"
import Vuex from "vuex"

Vue.use(Vuex)

const createStore = () => {
  return new Vuex.Store({
    state: () => ({
      counter: 0,
      show: false,
    }),
    mutations: {
      increment (state) {
        state.counter++
      },
      display (state) {
        state.show = !state.show
        console.log(state.show)
      }
    }
  })
}

export default createStore
