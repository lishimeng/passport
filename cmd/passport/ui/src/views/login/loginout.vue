<template>
</template>

<script  setup lang="ts" name="logout">
import {onMounted} from "vue";
import {clearTokenApi} from "/@/api/login";
import {Local, Session} from "/@/utils/storage";
import {useRoute, useRouter} from "vue-router";
const router = useRouter();
const route = useRoute();
onMounted(() => {
  var params=route.query
  var openUrl=''
  if(params&&params.path){
    openUrl="./#/login?path="+encodeURIComponent(params.path)
  }else{
    openUrl="./#/login"
  }
  clearTokenApi({}).then(res=>{
    Session.clear()
    Local.clear()
    window.location.replace(openUrl);
  })
});
</script>

<style scoped>

</style>