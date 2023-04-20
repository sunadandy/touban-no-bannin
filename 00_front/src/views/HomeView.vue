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
            <tr v-for="(touban, key) in toubans" :key="touban.id">
              <td class="touban-name">{{ touban.name }}</td>
              <td class="touban-owner">{{ touban.owner }}</td>
              <td><v-btn flat :rounded="0" color="success" @click="Show(key)">覗く</v-btn></td>
              <td><v-btn flat :rounded="0" color="success" @click="Edit(key)">編集</v-btn></td>
              <td><v-btn flat :rounded="0" color="success" @click="Delete(key)">削除</v-btn></td>
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
  data: function() {
    return {
      toubans: []
    }
  },
  mounted: function(){
    this.toubans = [
        {"id": 0, "name": "test1", "owner": "yamada"},
        {"id": 1, "name": "test2", "owner": "tanaka"},
    ]
  },
  methods: {
    Show: function(key) {
      this.$router.push({name: "show", params: {id: key}})
    },
    Edit: function(key) {
      this.$router.push({name: "edit", params: {id: key}})
    },
    Delete: function(key) {
      console.log(key)
      // this.$router.push({name: "RoomView", params: {id: key}})
    },
    Create: function() {
      this.$router.push("/create")
    },
  }
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