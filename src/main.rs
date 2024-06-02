// #![allow(non_upper_case_globals)]
// #![allow(non_camel_case_types)]
// #![allow(non_snake_case)]
#![allow(warnings, clippy::all)]
mod bind {
    include!(concat!(env!("OUT_DIR"), "/bindings.rs"));
}

fn main() {
    let n = 10;
    unsafe {
        let result = bind::Fibonacci(n);
        println!("Fibonacci({}) = {}", n, result);
    }

    println!("run snark verify");

    // unsafe{
        // bind::Snark_run();
    // }
}
