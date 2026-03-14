<script setup lang="ts">
defineProps<{
    name: string
    tagline: string
}>()
</script>

<template>
    <section class="min-h-screen flex flex-col justify-center px-8 md:px-20">
        <p class="text-sm text-gray-400 tracking-widest uppercase mb-3">Hi, saya</p>
        <h1 class="text-5xl md:text-7xl font-bold text-white mb-4">{{ name }}</h1>
        <p class="text-xl md:text-2xl text-gray-400 mb-8 max-w-xl">{{ tagline }}</p>
        <div class="flex gap-4">
            <a href="#projects"
               class="px-6 py-3 bg-white text-black font-medium rounded-full hover:bg-gray-200 transition">
                Lihat Projects
            </a>
            <a href="#contact"
               class="px-6 py-3 border border-gray-600 text-white rounded-full hover:border-white transition">
                Kontak Saya
            </a>
        </div>
    </section>
</template>