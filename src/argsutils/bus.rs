use crate::errs::errors;

use std::env::current_dir;
use std::path::Path;

pub struct LogicBus {
    pub conf_path: String,
    pub function: String,
    pub verbosity: i8,
    pub paths: Vec<String>,
    pub help: bool,
    pub called_from: String,
    pub user_dir: String,
    pub diff: bool,
    pub diff_opts: String,
    pub recursive: bool,
}

impl LogicBus {
    fn handle_verbose(&mut self, v: &str) {
        self.verbosity += v.chars().filter(|c| c == &'v').count() as i8;
    }

    pub fn assign_flag(&mut self, r: &str) {
        let rg = &r.replace("-", "");

        match rg.as_str() {
            // super clean match syntax, Rust can be pretty nice
            "d" | "diff" => self.diff = true,
            "h" | "help" => self.help = true,
            "r" | "recursive" => self.recursive = true,
            "v" | "vv" | "vvv" => self.handle_verbose(rg),
            _ => errors::err_flag(r),
        }
    }

    pub fn assign_func(&mut self, r: &str) {
        self.function = r.to_owned();
    }

    pub fn add_dir_called_from(&mut self) {
        let cd = current_dir().unwrap();
        self.called_from = cd.to_str().unwrap().to_owned();
    }

    pub fn add_path(&mut self, s: &str) {
        let p = Path::new(s).canonicalize().unwrap();
        let ps = p.to_str();
        self.paths.push(ps.unwrap().to_owned());
    }

    pub fn add_diff_opt(&mut self, s: &str) {
        if self.diff_opts.len() == 0 {
            self.diff_opts = s.to_owned();
        } else {
            errors::diff_exists_err();
        }
    }

    pub fn new() -> LogicBus {
        return LogicBus {
            conf_path: String::from(""),
            function: String::from(""),
            verbosity: 0,
            paths: vec![],
            help: false,
            called_from: String::from(""),
            user_dir: String::from(""),
            diff: false,
            diff_opts: String::from(""),
            recursive: false,
        };
    }
}
