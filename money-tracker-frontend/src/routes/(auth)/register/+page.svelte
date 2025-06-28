<script lang="ts">
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

	function enhance() {
		registerLoading = true;
		return async ({ result }: any) => {
			console.log(result);

			try {
				if (result.type === 'success') {
				}
			} catch (error) {
				console.error('error registerEnhance', error);
			} finally {
				registerLoading = false;
			}
		};
	}
</script>

<div class="mx-auto my-12 min-h-screen max-w-screen-lg border-0 md:border">
	<div class="grid grid-cols-1 items-center md:grid-cols-2">
		<div
			class="hidden h-full items-center justify-center bg-gray-800 py-12 md:flex"
			data-aos="fade-up"
			data-aos-duration="600"
		>
			<div class="space-y-6 px-8 text-center text-white">
				<div class="space-y-4">
					<h1 class="text-4xl font-bold">Money Tracker</h1>
					<div class="mx-auto h-1 w-16 rounded bg-white"></div>
					<p class="text-xl leading-relaxed text-gray-300">
						Take control of your finances<br />
						and build the future you deserve
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
					<p class="mb-3 text-sm text-gray-400">Already have an account?</p>
					<a
						href="/login"
						class="inline-block rounded-lg border border-white px-6 py-2 text-sm font-medium text-white transition-all duration-200 hover:bg-white hover:text-gray-800"
					>
						Sign In
					</a>
				</div>
			</div>
		</div>
		<div class="px-3 py-8" data-aos="fade-up" data-aos-delay="200" data-aos-duration="600">
			<Card>
				<CardHeader>
					<CardTitle class="text-2xl">Start Your Money Journey!</CardTitle>
					<CardDescription
						>Turn your financial dreams into reality, one step at a time</CardDescription
					>
				</CardHeader>
				<CardContent>
					<RegisterForm {username} {email} {password} {confirm_password} {enhance} />
				</CardContent>
			</Card>
		</div>
	</div>
</div>
