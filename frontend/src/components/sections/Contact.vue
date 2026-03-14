<script setup lang="ts">
import { reactive, ref } from 'vue'

const form = reactive({ name: '', email: '', message: '' })
const sending = ref(false)
const sent = ref(false)
const sendError = ref<string | null>(null)

// Nanti di-connect ke Go backend / Formspree
const handleSubmit = async () => {
    sending.value = true
    sendError.value = null
    try {
        // TODO: ganti URL dengan endpoint Go lo
        // await axios.post('http://localhost:8080/api/contact', form)
        await new Promise(r => setTimeout(r, 1000)) // simulasi dulu
        sent.value = true
        form.name = form.email = form.message = ''
    } catch {
        sendError.value = 'Gagal mengirim pesan, coba lagi.'
    } finally {
        sending.value = false
    }
}
</script>

<template>
    <section id="contact" class="py-24 px-8 md:px-20 border-t border-gray-800">
        <p class="text-sm text-gray-400 tracking-widest uppercase mb-3">Contact</p>
        <h2 class="text-3xl font-bold text-white mb-2">Hubungi Saya</h2>
        <p class="text-gray-400 mb-10">Tertarik untuk berkolaborasi? Drop a message.</p>

        <form @submit.prevent="handleSubmit" class="max-w-lg flex flex-col gap-4">
            <input
                v-model="form.name"
                type="text"
                placeholder="Nama"
                required
                class="bg-gray-900 border border-gray-700 rounded-lg px-4 py-3 text-white placeholder-gray-500 focus:outline-none focus:border-gray-400 transition"
            />
            <input
                v-model="form.email"
                type="email"
                placeholder="Email"
                required
                class="bg-gray-900 border border-gray-700 rounded-lg px-4 py-3 text-white placeholder-gray-500 focus:outline-none focus:border-gray-400 transition"
            />
            <textarea
                v-model="form.message"
                placeholder="Pesan"
                rows="5"
                required
                class="bg-gray-900 border border-gray-700 rounded-lg px-4 py-3 text-white placeholder-gray-500 focus:outline-none focus:border-gray-400 transition resize-none"
            />
            <button
                type="submit"
                :disabled="sending"
                class="px-6 py-3 bg-white text-black font-medium rounded-full hover:bg-gray-200 transition disabled:opacity-50"
            >
                {{ sending ? 'Mengirim...' : 'Kirim Pesan' }}
            </button>
            <p v-if="sent" class="text-green-400 text-sm">Pesan berhasil dikirim!</p>
            <p v-if="sendError" class="text-red-400 text-sm">{{ sendError }}</p>
        </form>
    </section>
</template>