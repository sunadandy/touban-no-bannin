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
            v-for="(m, index) in member" :key="m.order_number"
            v-bind:class="m.name == owner ? 'coloring' : ''"
          >
            <td>{{ m.order_number }}</td>
            <td>{{ m.name }}</td>
            <td v-if="editingCell.row!==index || editingCell.key!=='next'" @click="edit(index, m.name, 'next', m.next)">{{ m.next }}</td>
            <td v-else>
              <v-text-field
                v-model="editingCell.value"
                @blur="save(index)"
                @keydown.enter="save(index)"
                class="w-100"
                :rules="[dateRules]"
                autofocus
              ></v-text-field>
            </td>
            <td v-if="editingCell.row!==index || editingCell.key!=='last'" @click="edit(index, m.name, 'last', m.last)">{{ m.last }}</td>
            <td v-else>
              <v-text-field
                v-model="editingCell.value"
                @blur="save(index)"
                @keydown.enter="save(index)"
                class="w-100"
                :rules="[dateRules]"
                autofocus
                ></v-text-field>
            </td>
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
      message: "",
      editingCell: {row:-1, name:"", key:"", value: ""}, //編集中のセルの管理
      dateRules: (value) => {
        const pattern = /^[0-9]{4}-(0[1-9]|1[0-2])-(0[1-9]|[12][0-9]|3[01])$/;
        return pattern.test(value) || "yyyy-mm-dd形式で入力してください";
      },
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
  },
  methods:{
    edit(index, name, key, value){
      // 編集モードに入る際にもとの値を保持
      this.editingCell = {row: index, name:name, key:key, value:value};
    },
    save(index){
      // 保存された値が修正前と異なる場合はデータを更新して、DBも更新
      if((this.member[index][this.editingCell.key] != this.editingCell.value) && this.dateRules(this.editingCell.value) == true){
        // 値の更新
        this.member[index][this.editingCell.key] = this.editingCell.value
        // DB更新（変更した行だけputすればいいのだが、APIがメンバーリスト全体になっているので・・・・）
        this.axios
          .put("/member", this.member)
          .then((response) => {
            console.log(response);
          })
          .catch((error) => console.log(error));
      }else{
        // 元の値に戻す
        this.editingCell.value = this.member[index][this.editingCell.key]
      }

      // セル初期化
      this.editingCell = {row: -1, name:"", key:"", value:""};
    },
  },
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