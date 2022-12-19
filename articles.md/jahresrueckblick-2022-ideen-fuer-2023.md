---
title: 'Jahresrückblick 2022 und Ideen für 2023'
subtitle: 'Was war, was (vielleicht) wird.'
author: 'Patrick Bucher'
date: 2022-12-13T21:50:00
lang: de
---

Letztes Jahr habe ich einen
[Jahresrückblick](./2021-12-17-jahresrueckblick-2021.html) geschrieben und kurz
darauf einige [Ideen für 2022](./2021-12-31-ideen-fuer-2022.html)
aufgeschrieben. Was ist ‒ ersteres wiederholend ‒ aus letzterem geworden?

# Ideen umgesetzt?

Vorweg: die drei Themen bzw. Programmiersprachen, die ich mir notiert habe ‒
Elixir, Go und JavaScript ‒ haben mich alle beschäftigt. Oberflächlich
betrachtet habe ich also 2022 das gemacht, was ich mir vorgenommen hatte. Doch
wie sieht es bei genauerer Betrachtung aus?

## Elixir

Tatsächlich habe ich mich von Januar bis April immer wieder mit Elixir befasst,
wie die Commits im entsprechenden
[Repository](https://github.com/patrickbucher/elixir-basics) zeigen. Das Buch
[Elixir in
Action](https://www.manning.com/books/elixir-in-action-second-edition) habe ich
jedoch nur etwa zu zwei Dritteln durchgearbeitet. Ich bin bei den eigentlich
interessanten Teilen (Fehlerbehandlung) hängen geblieben.

Die 
[Elixir-Implementierung meines
SuperLeague-Programms](https://github.com/patrickbucher/superleague-polyglot/tree/master/superleague-elixir)
habe ich überarbeitet und dabei deutlich verbessern können. Soweit das Positive.

Mit Phoenix und dem Neuschreiben meiner
[Reversi-Simulation](https://github.com/patrickbucher/revergo) in Elixir ist es
jedoch nichts geworden.

Das einzige grössere Beispiel war eine nebenläufige
[Primzahlensuche](https://stackoverflow.com/q/71456020/6763074), die dann nicht
so recht skalieren wollte. Mit Go habe ich das Problem hingegen auf Anhieb lösen
können. Das dürfte dann sogleich der Ausstiegspunkt bzw. der Umstiegspunkt auf
Go gewesen sein.

## Go

Mit Generics habe ich mich nur im kleinen Rahmen befasst. Im Zusammenhang mit
der funktionalen Programmierung konnte ich Generics für eine kleine
[filter/map/reduce-Library](https://github.com/seantis/go-functools) nutzen.
Auch habe ich mir [Monaden in Go](./2022-11-05-towards-monads-in-go.html)
angeschaut, jedoch ohne generische Sprachkonstrukte.

Auf der Arbeit habe ich ein Memory-leakendes Logging-Tool (in Python
implementiert) durch ein einfacheres Tool in Go ersetzt. Das habe ich in Etappen
gemacht; mit dem Ergebnis bin ich recht zufrieden. Ich habe es schon mehrmals
verbessert und erweitert, was recht einfach ging.

Mit [Cloud Native
Go](https://www.oreilly.com/library/view/cloud-native-go/9781492076322/) und [12
Factor Apps](https://12factor.net/) habe ich mich nur kurz befasst, da ich für
das Cloud-Modul an der Berufsschule kaum über die Grundlagen herausgekommen bin.

Auch [gin](https://github.com/gin-gonic/gin) habe ich mir nicht angeschaut,
wobei ich ohnehin eher zu [Fiber](https://gofiber.io/) oder
[Gorilla](https://www.gorillatoolkit.org/) tendieren würde.

An privaten Projekten sind doch einige kleinere Sachen in Go zusammengekommen:

- [checklinks](https://github.com/patrickbucher/checklinks): ein
  Kommandozeilenwerkzeug, der die Links einer Webseite überprüft, um tote Links
  finden zu können
- [openbsd_autoinstall](https://github.com/patrickbucher/openbsd_autoinstall):
  ein minimalistischer HTTP-Server, der eine `install.conf` für OpenBSD liefert,
  was etwa für das Bauen von Images mit [Packer](https://www.packer.io/)
  sinnvoll sein kann
- [dfdegoregexp](https://github.com/patrickbucher/dfdegoregexp): eine kleine
  Einführung in reguläre Ausdrücke mit Go, für das deutsche Debianforum
  geschrieben
- [meow](https://github.com/patrickbucher/meow): ein kleines Monitoring-Tool als
  Anschauungsmaterial für den Berufsschulunterricht
- [huffman](https://github.com/patrickbucher/huffman): Textkompression mit
  Huffman-Bäumen als Lernprojekt

Für den Berufsschulunterricht habe ich auch eine Reihe von
[Videos](https://www.youtube.com/@m346pb/videos) mit Go-Bezug aufgenommen.

Immerhin…

## JavaScript

Ab Februar habe ich an der Berufsschule ein Praxismodul zum Thema
Web-Entwicklung unterrichtet. Hierfür habe ich mich wieder einmal etwas mit
JavaScript befasst. Dabei habe ich eher [Node.js](https://nodejs.org/en/) als
[Deno](https://deno.land/) verwendet, obwohl mir letzteres besser gefällt. Es
wurde jedoch November, bis ich mein SuperLeague-Programm mit
[JavaScript/Deno](https://github.com/patrickbucher/superleague-polyglot/tree/master/superleague-javascript) geschrieben habe.

Neben kleineren Programmierbeispielen für den Unterricht habe ich das [Game of
Life](https://github.com/patrickbucher/js-game-of-life) in JavaScript
geschrieben. Ein Überraschungserfolg wurde mein kleines
[Jass-Spiel](https://github.com/patrickbucher/jassete), das auch vom
[Schweizer
Jassverzeichnis](https://jassverzeichnis.ch/online-jass-wettspiel-im-jass-stuebli/)
verlinkt worden ist, ohne dass ich es dort gemeldet hätte. Man braucht nur
eine genügend schmale Nische zu finden…

Zum Thema JavaScript habe ich auch einige
[Videos](https://www.youtube.com/@ipt6web-entwicklung264/videos) für den
Berufsschulunterricht aufgenommen. Beruflich habe ich JavaScript sonst kaum
verwendet.

---

Soviel zu den gesteckten Zielen.

# Und sonst so?

Es muss so im April gewesen sein, als ich Elixir erneut habe fallen lassen. Ich
muss mich wohl etwas mit Clojure beschäftigt haben, wovon die entsprechende
[SuperLeague-Implementierung](https://github.com/patrickbucher/superleague-polyglot/tree/master/superleague-clojure)
zeugt. (Das gleiche habe ich noch in
[C](https://github.com/patrickbucher/superleague-polyglot/tree/master/superleague-c)
und in
[Racket](https://github.com/patrickbucher/superleague-polyglot/tree/master/superleague-racket)
gemacht, doch von letzterem später mehr…)

Einige neue Bücher auf meinem Regal zeugen auch von dieser Beschäftigung mit
Clojure. Statt eines gründlich durchzuarbeiten, habe ich verschiedene Bücher
angelesen. So richtig vorwärtsgekommen bin ich dabei nicht. Frustriert von
meiner fehlenden Konzentration ‒ und im Hinblick auf ein zweites Halbjahr mit
nur sehr wenig Freizeit, habe ich dann Clojure wieder fallen lassen, jedoch
Scheme wieder aufgenommen.

## SICP: der zweite Versuch

Ich habe mir folgendes überlegt: Wenn ich
[SICP](https://mitpress.mit.edu/9780262510875/structure-and-interpretation-of-computer-programs/)
durcharbeite, profitiere ich dabei sicherlich auch für Clojure, sowie für andere
(funktionale) Programmiersprachen. Und wenn ich täglich daran arbeite ‒ und sei
es auch nur eine gelesene Seite oder eine angefangene Übung ‒ bleibe ich
sicherlich nicht stehen. Gerade das zweite Halbjahr würde mir die Vorbereitung
des Berufsschulunterrichts sehr viel Aufwand bereiten. So bleibe ich zumindest
nicht stehen in meiner Beschäftigung mit der funktionalen Programmierung.

So schaue ich nun auf über vier Monate [täglicher
Beschäftigung](https://github.com/patrickbucher/sicp/blob/master/diary.md) mit
SICP zurück. Sogar am [Tag meines
Umzugs](https://github.com/patrickbucher/sicp/blob/master/diary.md#2022-10-04-tu)
konnte ich eine Übung lösen, die ich dann per Smartphone-Hotspot auf GitHub
gepusht habe. Ja, ich habe es dieses mal durchgezogen!

Gelernt habe ich dabei so einiges. Einige Frustrationen konnte ich
gewinnbringend überwinden, in dem ich etwa auf Racket ausweichen musste, um die
Beispiele mit der Bildverarbeitung im zweiten Kapitel testen zu können. (Aus
dieser Beschäftigung stammte auch die SuperLeague-Implementierung.)

Kapitel 2 könnte ich noch dieses Jahr beenden. Kapitel 3 möchte ich sicherlich
auch noch durcharbeiten. Kapitel 4 und 5 hingegen könnten warten, und ich könnte
mein gewonnenes Wissen über die funktionale Programmierung vielleicht einmal
praktisch anwenden.

---

Für die Berufsschule habe ich mich noch etwas mit Packer befasst, um eine
Ubuntu-VM für den Unterricht automatisch bauen zu können. Packer war auch auf
der Arbeit ein Thema. Weiter habe ich mich wieder einmal etwas mit Redis und
MinIO befasst, auch das für die Berufsschule. Meine Shell-Skripts für die
[Gitea-Administration](https://github.com/patrickbucher?page=1&tab=repositories)
habe ich in Python umgeschrieben und erweitert, was sich durchaus gelohnt hat.

Auf der Arbeit habe ich keine grösseren Würfe zu verzeichnen. Ich habe vor mich
hingewerkelt und das eine oder andere verbessert. Ein Werkzeug zur Auslosung
interner Restore-Tests hat mich zur kurzen Beschäftigung mit
[GraphQL](https://graphql.org/) gebracht. Auch [Podman](https://podman.io/) war
kurz ein Thema. Mit Ruby und Vagrant hatte ich mich auch kurz beschäftigt,
jedoch zu wenig nachhaltig. Ansible ist auch wieder unter den Tisch gefallen,
Kubernetes genauso, wobei ich mir letzteres gar nicht vorgenommen hatte.

Mit zwei Freunden halte ich einen unregelmässizen Lesezirkel zu Gerald Weinbergs
_Psychology of Computer Programming_ ab. Das Buch ist ein wahrer Schatz, und zu
jedem Kapitel fallen jedem Teilnehmer verschiedenste bestätigende Beispiele ein.
Wir sind aber in einem halben Jahr nicht ganz durchgekommen. Ich freue mich
aber, den Lesezirkel 2023 weiterzuführen; auch mit einem weiteren Buch.

Sonst gibt es von 2022 nicht mehr viel zu berichten. Gelesen habe ich nicht sehr
viel. Das [Cryptonomicon](https://www.nealstephenson.com/cryptonomicon.html)
zieht sich seit Sommer hin. Wenigstens wäre nach dem Umzug meine Bibliothek
wieder einigermassen aufgeräumt. Den Kindle habe ich wieder aktiviert, mir
darauf aber v.a. leichte Lektüre zugeführt. Fremdsprachen waren kein Thema
dieses Jahr. Immerhin lese ich das meiste auf Englisch.

Zum Jahresende hat mir dann der [Adventskalender 2022 vom
Debianforum](https://wiki.debianforum.de/Adventskalender_2022) noch etwas Arbeit
und noch viel mehr Freude beschert. Zwei meiner «Türchen» kamen aus dem Umfeld
des Berufsschulunterrichts (Redis und ein noch geheimes); die anderen beiden von
der Arbeit (`spiped`) bzw. aus der Freizeit (Huffman-Codierung).

Das Unterrichten mit Vor- und Nachbereitung hat in der zweiten Jahreshälfte
nicht nur zwei volle Tage, sondern auch den grössten Teil des Wochenendes
eingenommen. Im Frühling habe ich nur halbsoviele Lektionen, bei gleicher
Bezahlung. (Die Pensen werden auf das ganze Schuljahr hochgerechnet; d.h. ich
hatte jetzt ein Semester lang etwas über 120% gearbeitet und kann Ende Januar
auf ca. 80% reduzieren.) Das gröbste habe ich überstanden, und 2023 dürfte es
etwas mehr Luft geben.

# Ideen für 2023

Was soll ich also mit der Zeit anfangen?

Bis Ende April stehen einige Servermigrationen an; d.h. nicht nur VMs, sondern
auch ein physischer im Serverraum. Weiter sollte ich mich mit
Container-Registries befassen und vielleicht Podman noch etwas genauer
anschauen. Kubernetes wäre auch wieder einmal interessant, zumal ich es bisher
v.a. aus der OpenShift-Perspektive kenne, und sich Kubernetes seit meiner
letzten Beschäftigung damit sicherlich weiterentwickelt hat. Auch mein
Python-Wissen sollte ich gelegentlich etwas vertiefen (z.B. mit asyncio oder der
ganzen Packaging-Thematik, ansonsten reizt mich daran derzeit wenig).

SICP wird mich sicherlich noch einige Monate beschäftigen, wenn nicht das ganze
Jahr. Ansonsten reizt mich derzeit so einiges: Elixir, Rust, Perl (kein
Tippfehler!), Clojure, Elm, Svelte und Racket. Viele dieser Sprachen und
Technologien (Perl ausgenommen) helfen dabei, Software besser zu entwickeln,
indem man auf unkontrollierte Seiteneffekte verzichtet.

Bei ben Betriebssystemen begnüge ich mich derzeit mit Linux. FreeBSD und OpenBSD
spielen bei mir derzeit (leider) keine Rolle. Mein Backup- und Dateiserver, auf
dem FreeBSD mit ZFS läuft, liegt immer noch in einer Umzugskiste. Im Frühling
soll ich zudem eine Linux-Einführung an der Berufsschule geben ‒ für die Lehrer,
nicht für die Schüler. Themen wie systemd, nftables und Netzwerke allgemein
sollte ich gelegentlich auch etwas vertiefen. Dazu kommen auch diverse
Cloud-Themen, u.a. auch Datenschutz und anbieterspezifische Themen.

Im zweiten Halbjahr steht dann die didaktische Grundausbildung an, die mich wohl
jeweils samstags beschäftigen wird. Geniesse ich also die grösstenteils freien
Wochenenden im ersten Halbjahr.
