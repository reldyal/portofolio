import { ref } from 'vue'
import axios from 'axios'

interface Repo {
    id: number
    name: string
    description: string | null
    html_url: string
    language: string | null
    stargazers_count: number
}

export function useGithub(username: string) {
    const repos = ref<Repo[]>([])
    const loading = ref(false)
    const error = ref<string | null>(null)

    const fetchRepos = async () => {
        loading.value = true
        try {
            const { data } = await axios.get<Repo[]>(
                `https://api.github.com/users/${username}/repos?sort=updated&per_page=10`
            )
            repos.value = data
        } catch (e) {
            error.value = 'Gagal fetch repos'
        } finally {
            loading.value = false
        }
    }

    return { repos, loading, error, fetchRepos }
}