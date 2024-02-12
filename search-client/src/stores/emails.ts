import {defineStore} from "pinia"
import {ref} from "vue"

interface IEmail {
    message_id: string,
    date: string,
    from: string,
    x_from: string,
    to: string,
    x_to: string,
    subject: string,
    cc: string,
    x_cc: string,
    bcc: string,
    x_bcc: string,
    content: string
}

export const useEmailsStore = defineStore('emails', () => {
    const emails = ref<IEmail[]>([])
    const searchWord = ref("")
    const currentEmail = ref<IEmail | null>(null)
    const totalPages = ref(0)
    const loading = ref<boolean>(false)
    const errorMsg = ref<string | null>(null)

    async function FetchEmails(term:string, from_email:number, max_emails: number) {
        loading.value = true
        try {
            const response = await fetch(`${import.meta.env.VITE_API_URL}/api/search_emails?term=${term}&from_email=${from_email}&max_emails=${max_emails}`, {
                method: 'GET',
                headers: {
                    'Content-Type': 'application/json',
                }
            })
            
            const data: {total_emails?: number; emails?: IEmail[]; error?:string} = await response.json()

            if (data.emails) {
                emails.value = data.emails
            }

            if (data.total_emails || data.total_emails == 0) {
                totalPages.value = data.total_emails
            }

            if (data.error) {
                errorMsg.value = data.error
            }
        } catch(error) {
            errorMsg.value = error as string
        } finally {
            loading.value = false
        }
    }

    const setCurrentEmail = (email: IEmail) => {
        currentEmail.value = email
    }

    const setSearchWord = (word: string) => {
        searchWord.value = word
    }

    return {emails, totalPages, loading, errorMsg, FetchEmails, currentEmail, setCurrentEmail, searchWord, setSearchWord}
})