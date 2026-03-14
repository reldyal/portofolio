<script setup lang="ts">
import Hero from '../components/sections/Hero.vue'
import Bio from '../components/sections/Bio.vue'
import Projects from '../components/sections/Projects.vue'
import Contact from '../components/sections/Contact.vue'
</script>

<template>
    <main class="bg-black min-h-screen">
        <Hero
            name="Varel Endy Al Faiz AKA Reldyal"
            tagline="Full-stack developer yang suka bikin hal yang berguna."
        />
        <Bio
            bio="Gw seorang developer dengan passion di backend dan sistem. Suka eksplorasi teknologi baru dan bikin project dari nol."
            :skills="['Go', 'Vue', 'TypeScript', 'PostgreSQL', 'Docker']"
        />
        <Projects username="reldyal" />
        <Contact />
    </main>
</template>