<script setup lang="ts">
defineProps<{
    bio: string
    skills: string[]
}>()
</script>

<template>
    <section id="about" class="py-24 px-8 md:px-20 border-t border-gray-800">
        <p class="text-sm text-gray-400 tracking-widest uppercase mb-3">About</p>
        <h2 class="text-3xl font-bold text-white mb-6">Tentang Saya</h2>
        <p class="text-gray-400 max-w-2xl leading-relaxed text-lg">{{ bio }}</p>

        <div class="flex flex-wrap gap-3 mt-8">
            <span
                v-for="skill in skills"
                :key="skill"
                class="px-4 py-1.5 bg-gray-800 text-gray-300 rounded-full text-sm"
            >
                {{ skill }}
            </span>
        </div>
    </section>
</template>