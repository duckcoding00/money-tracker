import { fail, type Actions } from '@sveltejs/kit';
import type { PageServerLoad } from './$types';

export const load = (async () => {
	return {};
}) satisfies PageServerLoad;

export const actions: Actions = {
	login: async ({ fetch, request, cookies }) => {
		const form = await request.formData();
		const email = form.get('email') as string;
		const password = form.get('password') as string;
		const failedCountStr = cookies.get('failed_count') || '0';
		let failed = parseInt(failedCountStr);

		const body = JSON.stringify({
			email,
			password
		});

		try {
			const response = await fetch('http://127.0.0.1:8080/api/v1/user/login', {
				method: 'POST',
				headers: {
					'Content-Type': 'application/json',
					Accept: 'application/json'
				},
				body: body
			});

			const result = await response.json();

			console.log(result);
			console.log(failed);
			if (!response.ok) {
				failed++;

				cookies.set('failed_count', failed.toString(), {
					path: '/',
					httpOnly: true,
					sameSite: 'lax',
					secure: false,
					maxAge: 60 * 60
				});

				if (failed > 3) {
					return fail(400, {
						failed
					});
				}
				return fail(response.status, {
					...result,
					failed
				});
			}
			cookies.delete('failed_count', { path: '/' });

			const accessToken = result.data?.access_token;
			const username = result.data?.username;

			if (accessToken) {
				cookies.set('access_token', accessToken, {
					path: '/',
					httpOnly: true,
					sameSite: 'lax',
					secure: false,
					maxAge: 60 * 60 * 24 * 7
				});
			}

			return {
				username: username
			};
		} catch (error) {
			return fail(500, { message: 'Server error occurred' });
		}
	},
	logout: async ({ cookies }) => {
		cookies.delete('access_token', {
			path: '/'
		});

		return {
			auth: false
		};
	}
};
