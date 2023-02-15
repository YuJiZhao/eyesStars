<template>
  <Sky />
</template>

<script setup lang="ts">
const contextStore = useContext();

useHead({
  title: contextStore.value.appTitle,
  meta: [
    { name: 'keywords', content: contextStore.value.metaKeyWord },
    { name: 'description', content: contextStore.value.metaDesc }
  ],
  link: [{ rel: "icon", type: "image/x-icon", href: contextStore.value.favicon }],
});
</script>

<style>
* {
  margin: 0;
  padding: 0;
}

body {
  background-color: #000;
  overflow: hidden;
}
</style>