import axios from 'axios'
import { createStore } from 'vuex'

export default createStore({
  state: {
    currenToubanTable:[],
    currentOrderTable:[],
  },
  getters: {
    GetToubanByID: (state) => (id) => {
      // 条件が一致しなかった場合、空配列がリターンされる
      return state.currenToubanTable.filter(item => item.id == id)
    },
    GetMemberByToubanId: (state) => (id) => {
      // 条件が一致しなかった場合、空配列がリターンされる
      return state.currentOrderTable.filter(item => item.touban_id == id)
    },
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
      axios.get("/member")
      .then(response => {
        context.commit('setCurrentOrderTable', response.data)
      })
      // .catch(error => console.log(error))
    }
  },
  modules: {
  }
})
