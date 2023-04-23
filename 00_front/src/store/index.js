import axios from 'axios'
import { createStore } from 'vuex'

export default createStore({
  state: {
    currenToubanTable:[],
    currentOrderTable:[],
  },
  getters: {
    GetOrderByID: (state) => (id) => {
      return state.currentOrderTable.filter(item => item.toubanId == id)
    }
  },
  mutations: {
    setCurrenToubanTable(state, data){
      state.currenToubanTable = data
    },
    setCurrentOrderTable(state, data){
      state.currentOrderTable = data
    },
  },
  actions: {
    fetchDB(context){
      axios.get("/touban")
      .then(response => {
        context.commit('setCurrenToubanTable', response.data)
      })
      // .catch(error => console.log(error))

      // [MEMO]必要以上にデータをフェッチしているので遅延に注意
      axios.get("/order")
      .then(response => {
        context.commit('setCurrentOrderTable', response.data)
      })
      // .catch(error => console.log(error))
    }
  },
  modules: {
  }
})
