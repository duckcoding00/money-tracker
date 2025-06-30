<script lang="ts">
	import type { PageData } from './$types';
	import type { user as User } from '$lib/types/types';
	import VerifyAccount from '$lib/components/custom/VerifyAccount.svelte';
	import { goto } from '$app/navigation';

	let { data }: { data: PageData } = $props();

	let user: User | undefined = $state(data.user);

	let otp = $state('');
	let username = $state(user?.Username);
	let loadingVerify = $state(false);
	let successMessage = $state('');
	let errorMessage = $state<string[]>([]);

	function verifyEnhance() {
		loadingVerify = true;
		return async ({ result }: any) => {
			try {
				if (result.type === 'success') {
					errorMessage = [];
					successMessage = 'SUCCESS VERIFY, LETS LOGIN AGAIN';
					setTimeout(() => {
						goto('/login', {
							invalidateAll: true
						});
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
				console.error('error loginEnhance', error);
				errorMessage = ['An unexpected error occurred. Please try again.'];
				successMessage = '';
			} finally {
				loadingVerify = false;
			}
		};
	}
</script>

{#if user?.IsActive == false}
	<VerifyAccount
		{username}
		{otp}
		loading={loadingVerify}
		successAlert={successMessage}
		errorAlert={errorMessage}
		enhace={verifyEnhance}
		action="?/verify"
	/>
{:else}
	<p>DASHBOARD PAGE</p>
{/if}
