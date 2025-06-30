import { jwtDecode } from 'jwt-decode';
import type { LayoutServerLoad } from './$types';
import { fail } from '@sveltejs/kit';
import type { user } from '$lib/types/types';

export const load = (async ({ fetch }) => {
	let auth = false;
	try {
		const response = await fetch('http://127.0.0.1:8080/api/v1/user');
		const result = await response.json();

		if (!response.ok) {
			return {
				auth: false,
				user: undefined
			};
		}

		const user: user = result.data;
		console.log(user);

		if (user) {
			auth = true;
		}
		return {
			user,
			auth
		};
	} catch (error) {
		return fail(500, { message: 'Failed to load user data' });
	}
}) satisfies LayoutServerLoad;
