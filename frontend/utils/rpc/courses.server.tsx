import { ServerPublicRPCRequest } from "./rpc.server";
import {messages as wrapper} from '@/messages/wrapper'
import {messages as courses_api} from '@/messages/courses/api'
import {messages as courses_helpers} from '@/messages/courses/helpers'

export async function GetCoursesTeasers(): Promise<courses_helpers.CourseTeaser[]> {
    const wrappedResponse = await ServerPublicRPCRequest(new wrapper.Wrapper({
        get_courses_teasers_request: new courses_api.GetCoursesTeasersRequest({})
    }), "getcoursesteasers");
    if (!wrappedResponse.has_get_courses_teasers_response) {
        throw new Error("no get_courses_teasers_response");
    }
    return wrappedResponse.get_courses_teasers_response.teasers;
}

export async function GetCourse(slug: string): Promise<courses_helpers.CourseDetails> {
    const wrappedResponse = await ServerPublicRPCRequest(new wrapper.Wrapper({
        get_course_request: new courses_api.GetCourseRequest({
            slug: slug,
        })
    }), `getcourse-${slug}`);
    if (!wrappedResponse.has_get_course_response) {
        throw new Error("no get_course_response");
    }
    return wrappedResponse.get_course_response.details;
}