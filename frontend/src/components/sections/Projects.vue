<script setup lang="ts">
import { onMounted } from 'vue'
import { useGithub } from '../../composables/useGithub'

const props = defineProps<{ username: string }>()
const { repos, loading, error, fetchRepos } = useGithub(props.username)

onMounted(() => fetchRepos())
</script>

<template>
    <section id="projects" class="py-24 px-8 md:px-20 border-t border-gray-800">
        <p class="text-sm text-gray-400 tracking-widest uppercase mb-3">Work</p>
        <h2 class="text-3xl font-bold text-white mb-10">Projects</h2>

        <div v-if="loading" class="text-gray-400">Loading...</div>
        <div v-else-if="error" class="text-red-400">{{ error }}</div>

        <div
            v-for="repo in repos"
            :key="repo.id"
            @click="() => window.open(repo.html_url, '_blank')"
            class="p-6 bg-gray-900 border border-gray-800 rounded-xl hover:border-gray-600 transition group cursor-pointer"
        >
            <div class="flex justify-between items-start mb-3">
                <h3 class="text-white font-semibold group-hover:text-gray-200">
                    {{ repo.name }}
                </h3>
                <span class="text-gray-500 text-xs">↗</span>
            </div>
            <p class="text-gray-400 text-sm leading-relaxed mb-4">
                {{ repo.description ?? 'No description.' }}
            </p>
            <div class="flex items-center gap-4 text-xs text-gray-500">
                <span v-if="repo.language">{{ repo.language }}</span>
                <span>⭐ {{ repo.stargazers_count }}</span>
            </div>
        </div>
    </section>
</template>