<script lang="ts">
	import { goto } from '$app/navigation';
	import LoginForm from '$lib/components/custom/LoginForm.svelte';
	import { Card } from '$lib/components/ui/card';
	import CardContent from '$lib/components/ui/card/card-content.svelte';
	import CardDescription from '$lib/components/ui/card/card-description.svelte';
	import CardHeader from '$lib/components/ui/card/card-header.svelte';
	import CardTitle from '$lib/components/ui/card/card-title.svelte';
	import type { PageData } from './$types';

	let { data }: { data: PageData } = $props();

	let email = $state('');
	let password = $state('');
	let loginLoading = $state(false);
	let errorMessage = $state<string[]>([]);
	let successMessage = $state('');
	let failedCredentials = $state(false);

	function enhance() {
		return async ({ result }: any) => {
			console.log(result);
			loginLoading = true;
			errorMessage = [];
			successMessage = '';

			try {
				if (result.type === 'success') {
					errorMessage = [];
					successMessage = 'Success Login, redirecting...';
					setTimeout(() => {
						goto('/dashboard', {
							invalidateAll: true
						});
					}, 2000);
				} else if (result.type === 'failure') {
					successMessage = '';
					let { details, message, failed } = result.data;

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

					if (failed > 3) {
						errors.push("Forgot your credentials? Don't worry, just click here!");
						failedCredentials = true;
					}

					errorMessage = errors;
				}
			} catch (error) {
				console.error('error loginEnhance', error);
				errorMessage = ['An unexpected error occurred. Please try again.'];
				successMessage = '';
			} finally {
				loginLoading = false;
			}
		};
	}
</script>

<div
	class="mx-auto my-12 min-h-screen max-w-screen-lg border-8 border-black bg-white shadow-[20px_20px_0px_0px_#000000]"
>
	<div class="grid min-h-screen grid-cols-1 items-center md:grid-cols-2">
		<div
			class="flex items-center justify-center border-r-8 border-black px-8 py-8"
			data-aos="fade-up"
			data-aos-duration="600"
		>
			<div class="w-full max-w-md">
				<Card
					class="-rotate-1 transform border-6 border-black bg-white shadow-[15px_15px_0px_0px_#666666]"
				>
					<CardHeader class="rotate-1 transform border-b-4 border-black bg-gray-800">
						<CardTitle
							class="-rotate-1 transform text-3xl font-black tracking-wider text-white uppercase"
							>WELCOME BACK!</CardTitle
						>
						<CardDescription class="text-sm font-bold tracking-wide text-gray-300 uppercase"
							>SIGN IN TO CONTINUE YOUR FINANCIAL JOURNEY</CardDescription
						>
					</CardHeader>
					<CardContent class="p-6">
						<LoginForm
							bind:email
							bind:password
							{failedCredentials}
							loading={loginLoading}
							{enhance}
							errorAlert={errorMessage}
							successAlert={successMessage}
							action="?/login"
						/>
					</CardContent>
				</Card>
			</div>
		</div>

		<div
			class="hidden h-full items-center justify-center border-black bg-gray-600 py-12 md:flex"
			data-aos="fade-up"
			data-aos-delay="200"
			data-aos-duration="600"
		>
			<div class="space-y-8 px-8 text-center text-white">
				<div class="space-y-6">
					<div
						class="rotate-2 transform border-4 border-white bg-black p-4 shadow-[10px_10px_0px_0px_#333333]"
					>
						<h1 class="text-4xl font-black tracking-widest text-white uppercase">MONEY TRACKER</h1>
					</div>
					<div
						class="-rotate-1 transform border-4 border-black bg-white p-4 shadow-[8px_8px_0px_0px_#333333]"
					>
						<p class="text-xl font-bold tracking-wide text-black uppercase">
							WELCOME BACK!<br />
							CONTINUE YOUR FINANCIAL JOURNEY
						</p>
					</div>
				</div>

				<div class="space-y-4">
					<div
						class="flex rotate-1 transform items-center justify-center gap-4 border-4 border-white bg-gray-800 p-3 shadow-[6px_6px_0px_0px_#333333]"
					>
						<div class="h-3 w-3 border-2 border-black bg-white"></div>
						<span class="text-sm font-bold tracking-wide uppercase">TRACK EXPENSES EASILY</span>
					</div>
					<div
						class="flex -rotate-1 transform items-center justify-center gap-4 border-4 border-black bg-white p-3 shadow-[6px_6px_0px_0px_#333333]"
					>
						<div class="h-3 w-3 border-2 border-white bg-black"></div>
						<span class="text-sm font-bold tracking-wide text-black uppercase"
							>SET FINANCIAL GOALS</span
						>
					</div>
					<div
						class="flex rotate-1 transform items-center justify-center gap-4 border-4 border-white bg-gray-800 p-3 shadow-[6px_6px_0px_0px_#333333]"
					>
						<div class="h-3 w-3 border-2 border-black bg-white"></div>
						<span class="text-sm font-bold tracking-wide uppercase">MONITOR YOUR PROGRESS</span>
					</div>
				</div>

				<div class="pt-6">
					<div
						class="rotate-2 transform border-4 border-white bg-black p-4 shadow-[8px_8px_0px_0px_#333333]"
					>
						<p class="mb-4 text-sm font-bold text-gray-300 uppercase">DON'T HAVE AN ACCOUNT YET?</p>
						<a
							href="/register"
							class="inline-block -rotate-1 transform border-4 border-black bg-white px-6 py-3 text-sm font-black tracking-wider text-black uppercase shadow-[6px_6px_0px_0px_#666666] transition-all duration-200 hover:translate-x-2 hover:translate-y-2 hover:rotate-0 hover:shadow-[3px_3px_0px_0px_#333333]"
						>
							🚀 SIGN UP
						</a>
					</div>
				</div>
			</div>
		</div>
	</div>
</div>
