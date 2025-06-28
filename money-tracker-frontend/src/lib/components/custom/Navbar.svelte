<script lang="ts">
	import AvatarFallback from '../ui/avatar/avatar-fallback.svelte';
	import AvatarImage from '../ui/avatar/avatar-image.svelte';
	import Avatar from '../ui/avatar/avatar.svelte';
	import Button from '../ui/button/button.svelte';
	import Input from '../ui/input/input.svelte';
	import ModalForm from './ModalForm.svelte';

	type props = {
		name?: string;
		status?: string;

		auth?: boolean;
	};

	let openModal = $state(false);
	let { name = 'joko', status = 'active', auth = false }: props = $props();
</script>

<div class="border border-black px-6 py-4">
	<div class="grid grid-cols-2">
		<div class="ml-4 flex items-center justify-start gap-6">
			<p class="text text-base font-black">Logo Brand</p>
			<div class="flex flex-1 flex-row space-x-2">
				<Input type="text" placeholder="Search.." class="flex-1 " />
				<Button variant="outline">Search</Button>
			</div>
		</div>
		<div class="mr-4 flex items-center justify-end gap-3">
			{#if auth}
				<Avatar
					class="h-10 w-10 ring-2 ring-gray-200 transition-all duration-200 hover:ring-gray-300"
				>
					<AvatarImage src="icon.jpg" alt={`${name} avatar`} />
					<AvatarFallback
						class="bg-gradient-to-br from-blue-500 to-purple-600 font-semibold text-white"
					>
						{name?.charAt(0).toUpperCase()}
					</AvatarFallback>
				</Avatar>
				<div class="flex min-w-0 flex-col justify-center">
					<p class="truncate text-sm leading-tight font-semibold text-gray-900">{name}</p>
					<p class="flex items-center gap-1 text-xs leading-tight text-gray-500 capitalize">
						<span
							class="h-2 w-2 rounded-full {status === 'active'
								? 'bg-green-500'
								: status === 'away'
									? 'bg-yellow-500'
									: 'bg-gray-400'}"
						></span>
						{status}
					</p>
				</div>
			{:else}
				<Button
					size="sm"
					variant="outline"
					class="border-gray-800 duration-100 hover:border-none hover:bg-gray-800 hover:text-white"
					onclick={() => {
						openModal = !openModal;
					}}>Sign In</Button
				>
			{/if}
		</div>
	</div>
</div>

<ModalForm bind:open={openModal} />
