<template>
</template>

<script  setup lang="ts" name="logout">
import {onMounted} from "vue";
import {clearTokenApi} from "/@/api/login";
import {Local, Session} from "/@/utils/storage";
import {useRouter} from "vue-router";
const router = useRouter();
onMounted(() => {
  clearTokenApi({}).then(res=>{
    Session.clear()
    Local.clear()
    window.location.replace("./#/login");
  })
});
</script>

<style scoped>

</style>