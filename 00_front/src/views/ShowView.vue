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
            v-for="m in member" :key="m.order"
            v-bind:class="m.owner ? 'coloring' : ''"
          >
            <td>{{ m.order }}</td>
            <td>{{ m.name }}</td>
            <td>{{ m.last }}</td>
            <td>{{ m.next }}</td>
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
      member: []
    }
  },
  mounted(){
    var urlPathId = this.$route.params.id
    this.member = this.$store.getters.GetOrderByID(urlPathId)
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