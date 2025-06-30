<script lang="ts">
	import Navbar from '$lib/components/custom/Navbar.svelte';
	import type { Snippet } from 'svelte';
	import type { PageData } from './$types';
	import type { user as User } from '$lib/types/types';
	import NavSide from '$lib/components/custom/NavSide.svelte';
	import { goto } from '$app/navigation';
	import { page } from '$app/state';

	let { children, data }: { children: Snippet; data: PageData } = $props();
	let auth = $state(data.auth);
	let user: User | undefined = $state(data.user);
	let loadingLogout = $state(false);
	function enhance() {
		loadingLogout = true;
		return async ({ result }: any) => {
			console.log(result);
			try {
				if (result.type === 'success') {
					let { auth } = result.data;
					auth = auth;
					setTimeout(() => {
						goto('/', {
							invalidateAll: true
						});
					}, 1000);
				}
			} catch (error) {
				console.error('error logout', error);
			} finally {
				loadingLogout = false;
			}
		};
	}

	$effect(() => {
		if (!auth && page.url.pathname !== '/') {
			goto('/');
		}

		if (auth && user && !user.IsActive) {
			goto('/dashboard');
			return;
		}
	});
</script>

{#if auth}
	{#if page.url.pathname != '/'}
		<Navbar {auth} name={user?.Username || 'Guest'} status={user?.IsActive || false} />
		<div class="grid grid-cols-12">
			<div class="col-span-2">
				<NavSide action="login?/logout" {enhance} isActive={user?.IsActive || false} />
			</div>
			<div class="col-span-10">
				{@render children()}
			</div>
		</div>
	{:else}
		{@render children()}
	{/if}
{:else if page.url.pathname === '/'}
	{@render children()}
{/if}
