<template>
  <div>
    <div class="box28">
      <span class="box-title">管理者からのメッセージ</span>
      <pre>{{ message }}</pre>
    </div>
    <div class="show-view">
      <v-table density="compact">
        <thead>
          <tr>
            <th class="text-center">
              実施順
            </th>
            <th class="text-center">
              メンバー
            </th>
            <th class="text-center">
              次回実施日
            </th>
            <th class="text-center">
              最終実施日
            </th>
          </tr>
        </thead>
        <tbody>
          <tr
            class="text-center"
            v-for="m in member" :key="m.order_number"
            v-bind:class="m.name == owner ? 'coloring' : ''"
          >
            <td>{{ m.order_number }}</td>
            <td>{{ m.name }}</td>
            <td>{{ m.next }}</td>
            <td>{{ m.last }}</td>
          </tr>
        </tbody>
      </v-table>
    </div>
    <div class="note">
      <p>・メンバーは担当順にソートされます。</p>
      <p>・オーナーは赤字で表示されます。</p>
    </div>
  </div>
</template>

<script>
export default {
  name: 'ShowView',
  data(){
    return {
      member: [],
      owner: "",
      message: ""
    }
  },
  created(){
    const toubanId = this.$route.params.id
    // オーナーの把握
    // [note]URL直接入力に対応するため、ページアクセス時にDBからデータを取得するようにする
    this.axios.get("/touban")
      .then(response => {
        const toubanInfos = response.data
        const toubanInfo = toubanInfos.filter(item => item.id == toubanId)
        this.owner = toubanInfo[0].owner  //GetToubanByIDがレコードフィルタ結果を配列で返してくる
        this.message = toubanInfo[0].message
      })
    this.axios.get("/member")
      .then(response => {
        const memberInfo = response.data
        const member = memberInfo.filter(item => item.touban_id == toubanId)
        // メンバーの取得
        this.member = member.sort((a, b) => a.order_number - b.order_number)
      })
  }
}
</script>

<style scoped>
.note{
  padding-bottom: 1em;
}
.v-table{
  margin: 0px 30px
}
.coloring{
  color: red
}
.box28 {
    position: relative;
    margin: 2em 0;
    padding: 25px 10px 7px;
    border: solid 2px #FFC107;
}
.box28 .box-title {
    position: absolute;
    display: inline-block;
    top: -2px;
    left: -2px;
    padding: 0 9px;
    height: 25px;
    line-height: 25px;
    font-size: 17px;
    background: #FFC107;
    color: #ffffff;
    font-weight: bold;
}
.box28 p {
    margin: 0; 
    padding: 0;
}
pre {
  font-family: Arial, sans-serif;
}
</style>