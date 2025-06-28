<script lang="ts">
	import LoginForm from '$lib/components/custom/LoginForm.svelte';
	import { Card } from '$lib/components/ui/card';
	import CardContent from '$lib/components/ui/card/card-content.svelte';
	import CardDescription from '$lib/components/ui/card/card-description.svelte';
	import CardHeader from '$lib/components/ui/card/card-header.svelte';
	import CardTitle from '$lib/components/ui/card/card-title.svelte';
	import type { PageData } from './$types';

	let { data }: { data: PageData } = $props();
</script>

<div class="mx-auto my-12 min-h-screen max-w-screen-lg border-0 md:border">
	<div class="grid min-h-screen grid-cols-1 items-center md:grid-cols-2">
		<div
			class="flex items-center justify-center px-3 py-8"
			data-aos="fade-up"
			data-aos-duration="600"
		>
			<Card class="w-full max-w-md">
				<CardHeader>
					<CardTitle class="text-2xl">Welcome Back!</CardTitle>
					<CardDescription>Sign in to continue your financial journey</CardDescription>
				</CardHeader>
				<CardContent>
					<LoginForm />
				</CardContent>
			</Card>
		</div>
		<div
			class="hidden h-full items-center justify-center bg-gray-800 py-12 md:flex"
			data-aos="fade-up"
			data-aos-delay="200"
			data-aos-duration="600"
		>
			<div class="space-y-6 px-8 text-center text-white">
				<div class="space-y-4">
					<h1 class="text-4xl font-bold">Money Tracker</h1>
					<div class="mx-auto h-1 w-16 rounded bg-white"></div>
					<p class="text-xl leading-relaxed text-gray-300">
						Welcome back!<br />
						Continue your financial journey
					</p>
				</div>
				<div class="space-y-3 text-gray-400">
					<div class="flex items-center justify-center gap-3">
						<div class="h-2 w-2 rounded-full bg-white"></div>
						<span class="text-sm">Track expenses easily</span>
					</div>
					<div class="flex items-center justify-center gap-3">
						<div class="h-2 w-2 rounded-full bg-white"></div>
						<span class="text-sm">Set financial goals</span>
					</div>
					<div class="flex items-center justify-center gap-3">
						<div class="h-2 w-2 rounded-full bg-white"></div>
						<span class="text-sm">Monitor your progress</span>
					</div>
				</div>
				<div class="pt-4">
					<p class="mb-3 text-sm text-gray-400">Don't have an account yet?</p>
					<a
						href="/register"
						class="inline-block rounded-lg border border-white px-6 py-2 text-sm font-medium text-white transition-all duration-200 hover:bg-white hover:text-gray-800"
					>
						Sign Up
					</a>
				</div>
			</div>
		</div>
	</div>
</div>
