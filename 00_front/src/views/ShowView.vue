<template>
  <div>
    <div class="abst">
      <p>・メンバーは担当順にソートされます。</p>
      <p>・オーナーは赤字で表示されます。</p>
    </div>
    <div class="owner-message">
      {{ message }}
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
  </div>
</template>

<script>
export default {
  name: 'ShowView',
  data(){
    return {
      member: [],
      owner: ""
    }
  },
  create(){
    // websocket代替
    location.reload()
  },
  mounted(){
    const toubanId = this.$route.params.id
    // オーナーの把握
    const toubanInfo = this.$store.getters.GetToubanByID(toubanId)
    this.owner = toubanInfo[0].owner  //GetToubanByIDがレコードフィルタ結果を配列で返してくる
    // メンバーの取得
    this.member = this.$store.getters.GetMemberByToubanId(toubanId)
  }
}
</script>

<style scoped>
.abst{
  padding-bottom: 1em;
}
.v-table{
  margin: 0px 30px
}
.coloring{
  color: red
}
</style>