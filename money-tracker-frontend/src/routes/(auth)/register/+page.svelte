<script lang="ts">
    import { goto } from '$app/navigation';
    import RegisterForm from '$lib/components/custom/RegisterForm.svelte';
    import { Card } from '$lib/components/ui/card';
    import CardContent from '$lib/components/ui/card/card-content.svelte';
    import CardDescription from '$lib/components/ui/card/card-description.svelte';
    import CardHeader from '$lib/components/ui/card/card-header.svelte';
    import CardTitle from '$lib/components/ui/card/card-title.svelte';
    import type { PageData } from './$types';

    let { data }: { data: PageData } = $props();

    let username = $state('');
    let email = $state('');
    let password = $state('');
    let confirm_password = $state('');
    let registerLoading = $state(false);
    let errorMessage = $state<string[]>([]);
    let successMessage = $state('');

    function enhance() {
        return async ({ result }: any) => {
            console.log(result);
            registerLoading = true;

            try {
                if (result.type === 'success') {
                    errorMessage = [];
                    successMessage = 'Success Register, redirecting to login page...';

                    setTimeout(() => {
                        goto('/login');
                    }, 2000);
                } else if (result.type === 'failure') {
                    successMessage = '';
                    let { details, message } = result.data;

                    const errors: string[] = [];

                    if (details) {
                        if (typeof details === 'object' && details !== null) {
                            Object.keys(details).forEach((field) => {
                                if (details[field]) {
                                    errors.push(details[field]);
                                }
                            });
                        } else if (typeof details === 'string') {
                            errors.push(details);
                        }
                    } else if (message) {
                        errors.push(message);
                    }

                    errorMessage = errors;
                }
            } catch (error) {
                console.error('error registerEnhance', error);
                errorMessage = ['An unexpected error occurred. Please try again.'];
                successMessage = '';
            } finally {
                registerLoading = false;
            }
        };
    }
</script>

<div class="mx-auto my-12 min-h-screen max-w-screen-lg border-8 border-black bg-white shadow-[20px_20px_0px_0px_#000000]">
    <div class="grid min-h-screen grid-cols-1 items-center md:grid-cols-2">
        <div
            class="hidden h-full items-center justify-center bg-gray-600 py-12 md:flex border-r-8 border-black"
            data-aos="fade-up"
            data-aos-duration="600"
        >
            <div class="space-y-8 px-8 text-center text-white">
                <div class="space-y-6">
                    <div class="transform rotate-2 bg-black border-4 border-white p-6 shadow-[12px_12px_0px_0px_#333333]">
                        <h1 class="text-5xl font-black text-white uppercase tracking-widest">MONEY TRACKER</h1>
                    </div>
                    <div class="transform -rotate-1 bg-white border-4 border-black p-4 shadow-[10px_10px_0px_0px_#333333]">
                        <p class="text-xl font-bold text-black uppercase tracking-wide">
                            TAKE CONTROL OF YOUR FINANCES<br />
                            BUILD THE FUTURE YOU DESERVE
                        </p>
                    </div>
                </div>

                <div class="space-y-4">
                    <div class="flex items-center justify-center gap-4 bg-gray-800 border-4 border-white p-4 shadow-[8px_8px_0px_0px_#333333] transform rotate-1">
                        <div class="h-4 w-4 bg-white border-2 border-black"></div>
                        <span class="text-sm font-bold uppercase tracking-wide">TRACK EXPENSES EASILY</span>
                    </div>
                    <div class="flex items-center justify-center gap-4 bg-white border-4 border-black p-4 shadow-[8px_8px_0px_0px_#333333] transform -rotate-1">
                        <div class="h-4 w-4 bg-black border-2 border-white"></div>
                        <span class="text-sm font-bold uppercase tracking-wide text-black">SET FINANCIAL GOALS</span>
                    </div>
                    <div class="flex items-center justify-center gap-4 bg-gray-800 border-4 border-white p-4 shadow-[8px_8px_0px_0px_#333333] transform rotate-1">
                        <div class="h-4 w-4 bg-white border-2 border-black"></div>
                        <span class="text-sm font-bold uppercase tracking-wide">MONITOR YOUR PROGRESS</span>
                    </div>
                </div>

                <div class="pt-6">
                    <div class="bg-black border-4 border-white p-6 shadow-[10px_10px_0px_0px_#333333] transform rotate-2">
                        <p class="mb-4 text-sm text-gray-300 font-bold uppercase tracking-wide">ALREADY HAVE AN ACCOUNT?</p>
                        <a
                            href="/login"
                            class="inline-block bg-white border-4 border-black px-8 py-3 text-sm font-black text-black uppercase tracking-wider shadow-[6px_6px_0px_0px_#666666] transition-all duration-200 hover:translate-x-2 hover:translate-y-2 hover:shadow-[3px_3px_0px_0px_#333333] transform -rotate-1 hover:rotate-0"
                        >
                            🔐 SIGN IN
                        </a>
                    </div>
                </div>
            </div>
        </div>

        <div class="px-8 py-12 bg-gray-200" data-aos="fade-up" data-aos-delay="200" data-aos-duration="600">
            <div class="w-full max-w-md mx-auto">
                <Card class="border-6 border-black bg-white shadow-[15px_15px_0px_0px_#666666] transform rotate-1">
                    <CardHeader class="bg-gray-800 border-b-4 border-black transform -rotate-1">
                        <CardTitle class="text-3xl font-black text-white uppercase tracking-wider transform rotate-1">
                            START YOUR MONEY JOURNEY!
                        </CardTitle>
                        <CardDescription class="text-gray-300 font-bold uppercase text-sm tracking-wide">
                            TURN YOUR FINANCIAL DREAMS INTO REALITY
                        </CardDescription>
                    </CardHeader>
                    <CardContent class="p-6">
                        <RegisterForm
                            bind:username
                            bind:email
                            bind:password
                            bind:confirm_password
                            loading={registerLoading}
                            {enhance}
                            errorAlert={errorMessage}
                            successAlert={successMessage}
                            action="?/register"
                        />
                    </CardContent>
                </Card>
            </div>
        </div>
    </div>
</div>
