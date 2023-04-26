<template>
  <div class="home">
    <div class="touban-list">
      <h2>現在作られている当番一覧</h2>
      <v-row class="touban-table" justify="center">
        <v-table>
          <thead>
            <tr>
              <th class="touban-name">当番名</th>
              <th class="touban-owner">オーナー</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="(touban, index) in this.$store.state.currenToubanTable" :key="index">
              <td class="touban-name">{{ touban.title }}</td>
              <td class="touban-owner">{{ touban.owner }}</td>
              <td><v-btn flat :rounded="0" color="success" @click="Show(touban.id)">覗く</v-btn></td>
              <td><v-btn flat :rounded="0" color="success" @click="Edit(touban.id)">編集</v-btn></td>
              <td><v-btn flat :rounded="0" color="success" @click="Delete(touban.id)">削除</v-btn></td>
            </tr>
          </tbody>
        </v-table>
      </v-row>
    </div>
    <div class="create">
      <v-btn
        :disabled="loading"
        :loading="loading"
        block
        class="text-none mb-4"
        color="indigo-darken-1"
        size="large"
        variant="flat"
        @click="Create(key)"
      >
      新規作成
      </v-btn>
    </div>
  </div>
</template>

<script>
export default {
  name: 'HomeView',
  methods: {
    Show: function(id) {
      this.$router.push({name: "show", params: {id: id}})
    },
    Edit: function(id) {
      this.$router.push({name: "edit", params: {id: id}})
    },
    Delete: function(id) {
      // 当番テーブルから該当するキーのレコードを削除
      this.axios.delete("/touban", {params: {id: id}})
      .then(
        this.$router.push("/home")
      )
      // 順番テーブルから当番IDが一致するレコードを削除
      this.axios.delete("/member", {params: {id: id}})
    },
    Create: function() {
      this.$router.push("/create")
    },
  },
}
</script>

<style scoped>
  h2 {
      text-align: center;
  }
  .home {
      border: 2px rgb(161, 158, 158) solid;
      margin: 10px 0px;
      padding: 20px 0px;
  }
  .touban-table {
    margin: auto 10px;
    margin-bottom: 20px;
  }
</style>