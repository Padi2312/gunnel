import type { PageLoad } from "./$types";
import { GetTunnelByID } from "../../lib/wailsjs/go/internal/Store";

export const prerender = true;
export const ssr = false;

export const load: PageLoad = async ({ url }) => {
    const id = url.searchParams.get("id") as string;
    const tunnel = await GetTunnelByID(id);
    return {
        tunnel
    }


}