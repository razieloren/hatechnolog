import { ServerPublicRPCRequest } from "./rpc.server";
import {messages as wrapper} from '@/messages/wrapper'
import {messages as content_api} from '@/messages/content/api'
import {messages as content_helpers} from '@/messages/content/helpers'

export async function GetPage(slug: string): Promise<content_helpers.ContentDetails> {
    const wrappedResponse = await ServerPublicRPCRequest(new wrapper.Wrapper({
        get_page_request: new content_api.GetPageRequest({
            slug: slug,
        })
    }), `getpage-${slug}`);
    if (!wrappedResponse.has_get_page_response) {
        throw new Error("no get_page_response");
    }
    return wrappedResponse.get_page_response.details;
}

export async function GetPostsTeasers(): Promise<content_helpers.ContentTeaser[]> {
    const wrappedResponse = await ServerPublicRPCRequest(new wrapper.Wrapper({
        get_posts_teasers_request: new content_api.GetPostsTeasersRequest({})
    }), "poststeasers");
    if (!wrappedResponse.has_get_posts_teasers_response) {
        throw new Error("no get_posts_teasers_response");
    }
    return wrappedResponse.get_posts_teasers_response.teasers;
}

export async function GetPost(slug: string): Promise<content_helpers.ContentDetails> {
    const wrappedResponse = await ServerPublicRPCRequest(new wrapper.Wrapper({
        get_post_request: new content_api.GetPostRequest({
            slug: slug,
        })
    }), `getpost-${slug}`);
    if (!wrappedResponse.has_get_post_response) {
        throw new Error("no get_post_response");
    }
    return wrappedResponse.get_post_response.details;
}

export async function GetCategoriesTeasers(): Promise<content_helpers.CategoryTeaser[]> {
    const wrappedResponse = await ServerPublicRPCRequest(new wrapper.Wrapper({
        get_categories_teasers_request: new content_api.GetCategoriesTeasersRequest({})
    }), "getcategoryteasers");
    if (!wrappedResponse.has_get_categories_teasers_response) {
        throw new Error("no get_categories_teasers_response");
    }
    return wrappedResponse.get_categories_teasers_response.teasers;
}

export async function GetCategory(slug: string): Promise<content_helpers.CategoryDetails> {
    const wrappedResponse = await ServerPublicRPCRequest(new wrapper.Wrapper({
        get_category_request: new content_api.GetCategoryRequest({
            slug: slug,
        })
    }), `getcategory-${slug}`);
    if (!wrappedResponse.has_get_category_response) {
        throw new Error("no get_category_response");
    }
    return wrappedResponse.get_category_response.details;
}