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
              最終実施日
            </th>
            <th class="text-center">
              次回実施日
            </th>
          </tr>
        </thead>
        <tbody>
          <tr
            v-for="item in items" :key="item.name"
            v-bind:class="item.owner ? 'coloring' : ''"
          >
            <td>{{ item.id }}</td>
            <td>{{ item.name }}</td>
            <td>{{ item.lasted }}</td>
            <td>{{ item.next }}</td>
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
    return{
      items: [
        {id: 0, name: "A", lasted: "2023/4/1", next: "2023/5/1", owner: false},
        {id: 1, name: "B", lasted: "2023/5/2", next: "2023/5/2", owner: true},
    ],
      message: "",
    }
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