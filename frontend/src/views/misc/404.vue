<template>
    <div class="404">
        <h1>404 Not Found</h1>
        <p>Halaman tidak ditemukan.</p>
        <RouterLink to="/">Kembali ke Home</RouterLink>
    </div>
</template>

<script setup lang="ts">
</script>

<style scoped>
.404 {
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    height: 100vh;
    gap: 1rem;
    text-align: center;
}

h1 {
    font-size: 6rem;
    margin: 0;
}
</style>